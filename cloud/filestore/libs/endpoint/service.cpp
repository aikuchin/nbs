#include "service.h"

#include "listener.h"

#include <cloud/filestore/libs/service/context.h>
#include <cloud/filestore/libs/service/endpoint.h>
#include <cloud/filestore/libs/service/request.h>

#include <cloud/storage/core/libs/common/error.h>
#include <cloud/storage/core/libs/coroutine/executor.h>
#include <cloud/storage/core/libs/diagnostics/logging.h>
#include <cloud/storage/core/libs/keyring/endpoints.h>

#include <util/generic/map.h>

namespace NCloud::NFileStore {

using namespace NThreading;

namespace {

////////////////////////////////////////////////////////////////////////////////

bool CompareRequests(
    const NProto::TEndpointConfig& left,
    const NProto::TEndpointConfig& right)
{
    return left.GetFileSystemId() == right.GetFileSystemId()
        && left.GetClientId() == right.GetClientId()
        && left.GetSocketPath() == right.GetSocketPath()
        && left.GetSessionRetryTimeout() == right.GetSessionRetryTimeout()
        && left.GetSessionPingTimeout() == right.GetSessionPingTimeout()
        && left.GetServiceEndpoint() == right.GetServiceEndpoint()
        && left.GetMountSeqNumber() == right.GetMountSeqNumber()
        && left.GetReadOnly() == right.GetReadOnly();
}

bool CompareRequests(
    const NProto::TStartEndpointRequest& left,
    const NProto::TStartEndpointRequest& right)
{
    return CompareRequests(left.GetEndpoint(), right.GetEndpoint());
}

bool SessionUpdateRequired(
    const NProto::TEndpointConfig& left,
    const NProto::TEndpointConfig& right)
{
    return left.GetMountSeqNumber() != right.GetMountSeqNumber()
        || left.GetReadOnly() != right.GetReadOnly();
}

////////////////////////////////////////////////////////////////////////////////

struct TEndpointInfo
{
    IEndpointPtr Endpoint;
    NProto::TEndpointConfig Config;
};

////////////////////////////////////////////////////////////////////////////////

struct TStartingEndpointState
{
    NProto::TStartEndpointRequest Request;
    TFuture<NProto::TStartEndpointResponse> Result;

    TStartingEndpointState(
            const NProto::TStartEndpointRequest& request,
            const TFuture<NProto::TStartEndpointResponse>& result)
        : Request(request)
        , Result(result)
    {}
};

////////////////////////////////////////////////////////////////////////////////

struct TStoppingEndpointState
{
    TFuture<NProto::TStopEndpointResponse> Result;

    TStoppingEndpointState(const TFuture<NProto::TStopEndpointResponse>& result)
        : Result(result)
    {}
};

////////////////////////////////////////////////////////////////////////////////

class TEndpointManager final
    : public IEndpointManager
    , public std::enable_shared_from_this<TEndpointManager>
{
private:
    const ILoggingServicePtr Logging;
    const IEndpointStoragePtr Storage;
    const IEndpointListenerPtr Listener;

    TExecutorPtr Executor = TExecutor::Create("SVC");
    TLog Log;

    THashMap<TString, TStartingEndpointState> StartingSockets;
    THashMap<TString, TStoppingEndpointState> StoppingSockets;

    TMap<TString, TEndpointInfo> Endpoints;

public:
    TEndpointManager(
            ILoggingServicePtr logging,
            IEndpointStoragePtr storage,
            IEndpointListenerPtr listener)
        : Logging(std::move(logging))
        , Storage(std::move(storage))
        , Listener(std::move(listener))
    {
        Log = Logging->CreateLog("NFS_SERVICE");
    }

    void Start() override
    {
        Executor->Start();
    }

    void Stop() override
    {
        Executor->Stop();

        TVector<TFuture<void>> futures;
        for (auto&& [_, endpoint]: Endpoints) {
            futures.push_back(endpoint.Endpoint->SuspendAsync());
        }
        WaitAll(futures).GetValueSync();
    }

#define FILESTORE_IMPLEMENT_METHOD(name, ...)                                  \
public:                                                                        \
    TFuture<NProto::T##name##Response> name(                                   \
        TCallContextPtr callContext,                                           \
        std::shared_ptr<NProto::T##name##Request> request) override            \
    {                                                                          \
        Y_UNUSED(callContext);                                                 \
        return Executor->Execute([this, request = std::move(request)] {        \
            return Do##name(*request);                                         \
        });                                                                    \
    }                                                                          \
private:                                                                       \
    NProto::T##name##Response Do##name(                                        \
        const NProto::T##name##Request& request);                              \
// FILESTORE_IMPLEMENT_METHOD

    FILESTORE_ENDPOINT_SERVICE(FILESTORE_IMPLEMENT_METHOD)

#undef FILESTORE_IMPLEMENT_METHOD

private:
    void AddStartingSocket(
        const TString& socketPath,
        const NProto::TStartEndpointRequest& request,
        const TFuture<NProto::TStartEndpointResponse>& result)
    {
        auto [it, inserted] = StartingSockets.emplace(
            std::piecewise_construct,
            std::forward_as_tuple(socketPath),
            std::forward_as_tuple(request, result));
        Y_ABORT_UNLESS(inserted);
    }

    void RemoveStartingSocket(const TString& socketPath)
    {
        StartingSockets.erase(socketPath);
    }

    void AddStoppingSocket(
        const TString& socketPath,
        const TFuture<NProto::TStopEndpointResponse>& result)
    {
        auto [it, inserted] = StoppingSockets.emplace(socketPath, result);
        Y_ABORT_UNLESS(inserted);
    }

    void RemoveStoppingSocket(const TString& socketPath)
    {
        StoppingSockets.erase(socketPath);
    }

    NProto::TError AddEndpointToStorage(
        const NProto::TStartEndpointRequest& request)
    {
        auto [data, error] = SerializeEndpoint(request);
        if (HasError(error)) {
            return error;
        }

        error = Storage->AddEndpoint(data).GetError();
        if (HasError(error)) {
            return error;
        }

        return {};
    }

    NProto::TError RemoveEndpointFromStorage(
        const NProto::TStopEndpointRequest& request)
    {
        const auto& socketPath = request.GetSocketPath();

        auto [ids, error] = Storage->GetEndpointIds();
        if (HasError(error)) {
            return error;
        }

        for (auto id: ids) {
            auto [data, error] = Storage->GetEndpoint(id);
            if (HasError(error)) {
                continue;
            }

            auto req = DeserializeEndpoint<NProto::TStartEndpointRequest>(data);
            if (req && req->GetEndpoint().GetSocketPath() == socketPath) {
                return Storage->RemoveEndpoint(id);
            }
        }

        return MakeError(E_INVALID_STATE, TStringBuilder()
            << "Couldn't find endpoint " << socketPath);
    }
};

////////////////////////////////////////////////////////////////////////////////

NProto::TStartEndpointResponse TEndpointManager::DoStartEndpoint(
    const NProto::TStartEndpointRequest& request)
{
    STORAGE_TRACE("StartEndpoint " << DumpMessage(request));

    const auto& config = request.GetEndpoint();
    const auto& socketPath = config.GetSocketPath();

    if (StoppingSockets.contains(socketPath)) {
        return TErrorResponse(E_REJECTED, TStringBuilder()
            << "endpoint " << socketPath.Quote()
            << " is stopping now");
    }

    if (const auto* p = StartingSockets.FindPtr(socketPath)) {
        if (!CompareRequests(request, p->Request)) {
            return TErrorResponse(E_REJECTED, TStringBuilder()
                << "endpoint " << socketPath.Quote()
                << " is starting now with other args");
        }

        auto future = p->Result;
        Executor->WaitFor(future);
        return future.GetValue();
    }

    if (auto* endpoint = Endpoints.FindPtr(socketPath)) {
        const auto& newConfig = request.GetEndpoint();
        const auto& oldConfig = endpoint->Config;
        if (CompareRequests(newConfig, oldConfig)) {
            return TErrorResponse(S_ALREADY, TStringBuilder()
                << "endpoint " << socketPath.Quote()
                << " already started");
        } else if (SessionUpdateRequired(newConfig, oldConfig)) {
            const auto& config = request.GetEndpoint();
            endpoint->Config.SetReadOnly(config.GetReadOnly());
            endpoint->Config.SetMountSeqNumber(config.GetMountSeqNumber());

            auto readOnly = config.GetReadOnly();
            auto mountSeqNumber = config.GetMountSeqNumber();

            auto future = endpoint->Endpoint->AlterAsync(
                readOnly,
                mountSeqNumber).Apply(
                [=] (const TFuture<NProto::TError>& future) {
                    NProto::TStartEndpointResponse response;
                    auto error = future.GetValue();
                    if (!HasError(error)) {
                        endpoint->Config.SetReadOnly(readOnly);
                        endpoint->Config.SetMountSeqNumber(mountSeqNumber);
                    }
                    response.MutableError()->CopyFrom(future.GetValue());
                    return response;
                });
            Executor->WaitFor(future);
            return future.GetValue();
        } else {
            return TErrorResponse(E_ARGUMENT, TStringBuilder()
                << "endpoint " << socketPath.Quote()
                << ": attempt to change non-modifiable parameters");
        }
    }

    auto endpoint = Listener->CreateEndpoint(config);
    auto future = endpoint->StartAsync().Apply(
        [=] (const TFuture<NProto::TError>& future) {
            NProto::TStartEndpointResponse response;
            response.MutableError()->CopyFrom(future.GetValue());

            return response;
        });

    AddStartingSocket(socketPath, request, future);
    Executor->WaitFor(future);
    RemoveStartingSocket(socketPath);

    const auto& response = future.GetValue();
    if (SUCCEEDED(response.GetError().GetCode())) {
        if (config.GetPersistent()) {
            AddEndpointToStorage(request);
        }

        Endpoints.emplace(socketPath, TEndpointInfo { endpoint, config });
    }

    return response;
}

NProto::TStopEndpointResponse TEndpointManager::DoStopEndpoint(
    const NProto::TStopEndpointRequest& request)
{
    STORAGE_TRACE("StopEndpoint " << DumpMessage(request));

    const auto& socketPath = request.GetSocketPath();

    if (StartingSockets.contains(socketPath)) {
        return TErrorResponse(E_REJECTED, TStringBuilder()
            << "endpoint " << socketPath.Quote()
            << " is starting now");
    }

    if (const auto* p = StoppingSockets.FindPtr(socketPath)) {
        auto future = p->Result;
        Executor->WaitFor(future);
        return future.GetValue();
    }

    auto it = Endpoints.find(request.GetSocketPath());
    if (it == Endpoints.end()) {
        return TErrorResponse(S_FALSE, TStringBuilder()
            << "endpoint " << socketPath.Quote()
            << " not found");
    }

    auto endpoint = it->second.Endpoint;
    auto future = endpoint->StopAsync().Apply(
        [] (const TFuture<void>&) {
            return NProto::TStopEndpointResponse{};
        });

    AddStoppingSocket(socketPath, future);
    Executor->WaitFor(future);
    RemoveStoppingSocket(socketPath);

    const auto& response = future.GetValue();
    if (SUCCEEDED(response.GetError().GetCode())) {
        auto config = it->second.Config;
        if (config.GetPersistent()) {
            RemoveEndpointFromStorage(request);
        }
    }

    Endpoints.erase(it);
    return future.GetValue();
}

NProto::TListEndpointsResponse TEndpointManager::DoListEndpoints(
    const NProto::TListEndpointsRequest& request)
{
    STORAGE_TRACE("ListEndpoints " << DumpMessage(request));

    NProto::TListEndpointsResponse response;
    for (const auto& [k, v]: Endpoints) {
        *response.AddEndpoints() = v.Config;
    }

    return response;
}

NProto::TKickEndpointResponse TEndpointManager::DoKickEndpoint(
    const NProto::TKickEndpointRequest& request)
{
    STORAGE_TRACE("KickEndpoint " << DumpMessage(request));

    auto requestOrError = Storage->GetEndpoint(request.GetKeyringId());
    if (HasError(requestOrError)) {
        return TErrorResponse(requestOrError.GetError());
    }

    auto startRequest = DeserializeEndpoint<NProto::TStartEndpointRequest>(
        requestOrError.GetResult());
    auto startResponse = DoStartEndpoint(*startRequest);

    NProto::TKickEndpointResponse response;
    response.MutableError()->CopyFrom(startResponse.GetError());

    return response;
}

NProto::TPingResponse TEndpointManager::DoPing(
    const NProto::TPingRequest& request)
{
    STORAGE_TRACE("Ping " << DumpMessage(request));

    return {};
}

////////////////////////////////////////////////////////////////////////////////

class TNullEndpointManager final
    : public IEndpointManager
{
    void Start() override
    {}

    void Stop() override
    {}

#define FILESTORE_IMPLEMENT_METHOD(name, ...)                                  \
    TFuture<NProto::T##name##Response> name(                                   \
        TCallContextPtr callContext,                                           \
        std::shared_ptr<NProto::T##name##Request> request) override            \
    {                                                                          \
        Y_UNUSED(callContext);                                                 \
        Y_UNUSED(request);                                                     \
        return MakeFuture<NProto::T##name##Response>();                        \
    }                                                                          \
// FILESTORE_IMPLEMENT_METHOD

    FILESTORE_ENDPOINT_SERVICE(FILESTORE_IMPLEMENT_METHOD)

#undef FILESTORE_IMPLEMENT_METHOD
};

}   // namespace

////////////////////////////////////////////////////////////////////////////////

IEndpointManagerPtr CreateEndpointManager(
    ILoggingServicePtr logging,
    IEndpointStoragePtr storage,
    IEndpointListenerPtr listener)
{
    return std::make_shared<TEndpointManager>(
        std::move(logging),
        std::move(storage),
        std::move(listener));
}

IEndpointManagerPtr CreateNullEndpointManager()
{
    return std::make_shared<TNullEndpointManager>();
}

}   // namespace NCloud::NFileStore
