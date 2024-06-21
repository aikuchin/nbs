#include "service_actor.h"

#include <cloud/filestore/libs/diagnostics/profile_log_events.h>
#include <cloud/filestore/libs/storage/api/tablet_proxy.h>
#include <cloud/filestore/libs/storage/tablet/model/verify.h>

#include <contrib/ydb/library/actors/core/actor_bootstrapped.h>

namespace NCloud::NFileStore::NStorage {

using namespace NActors;

using namespace NKikimr;

namespace {

////////////////////////////////////////////////////////////////////////////////

class TGetNodeAttrActor final: public TActorBootstrapped<TGetNodeAttrActor>
{
private:
    // Original request
    const TRequestInfoPtr RequestInfo;
    NProto::TGetNodeAttrRequest GetNodeAttrRequest;

    // Filesystem-specific params
    const TString LogTag;

    // Response data
    NProto::TGetNodeAttrResponse LeaderResponse;
    bool LeaderResponded = false;

    // Stats for reporting
    IRequestStatsPtr RequestStats;
    IProfileLogPtr ProfileLog;

    const bool MultiTabletForwardingEnabled;

public:
    TGetNodeAttrActor(
        TRequestInfoPtr requestInfo,
        NProto::TGetNodeAttrRequest getNodeAttrRequest,
        TString logTag,
        IRequestStatsPtr requestStats,
        IProfileLogPtr profileLog,
        bool multiTabletForwardingEnabled);

    void Bootstrap(const TActorContext& ctx);

private:
    STFUNC(StateWork);

    void GetNodeAttrInLeader(const TActorContext& ctx);

    void HandleGetNodeAttrResponse(
        const TEvService::TEvGetNodeAttrResponse::TPtr& ev,
        const TActorContext& ctx);

    void GetNodeAttrInFollower(const TActorContext& ctx);

    void HandlePoisonPill(
        const TEvents::TEvPoisonPill::TPtr& ev,
        const TActorContext& ctx);

    void ReplyAndDie(
        const TActorContext& ctx,
        NProto::TGetNodeAttrResponse followerResponse);
    void HandleError(const TActorContext& ctx, NProto::TError error);
};

////////////////////////////////////////////////////////////////////////////////

TGetNodeAttrActor::TGetNodeAttrActor(
        TRequestInfoPtr requestInfo,
        NProto::TGetNodeAttrRequest getNodeAttrRequest,
        TString logTag,
        IRequestStatsPtr requestStats,
        IProfileLogPtr profileLog,
        bool multiTabletForwardingEnabled)
    : RequestInfo(std::move(requestInfo))
    , GetNodeAttrRequest(std::move(getNodeAttrRequest))
    , LogTag(std::move(logTag))
    , RequestStats(std::move(requestStats))
    , ProfileLog(std::move(profileLog))
    , MultiTabletForwardingEnabled(multiTabletForwardingEnabled)
{
}

void TGetNodeAttrActor::Bootstrap(const TActorContext& ctx)
{
    GetNodeAttrInLeader(ctx);
    Become(&TThis::StateWork);
}

void TGetNodeAttrActor::GetNodeAttrInLeader(const TActorContext& ctx)
{
    LOG_DEBUG(
        ctx,
        TFileStoreComponents::SERVICE,
        "[%s] Executing GetNodeAttr in leader for %lu, %s",
        LogTag.c_str(),
        GetNodeAttrRequest.GetNodeId(),
        GetNodeAttrRequest.GetName().Quote().c_str());

    auto request = std::make_unique<TEvService::TEvGetNodeAttrRequest>();
    request->Record = GetNodeAttrRequest;

    // forward request through tablet proxy
    ctx.Send(MakeIndexTabletProxyServiceId(), request.release());
}

void TGetNodeAttrActor::GetNodeAttrInFollower(const TActorContext& ctx)
{
    LOG_DEBUG(
        ctx,
        TFileStoreComponents::SERVICE,
        "[%s] Executing GetNodeAttr in follower for %lu, %s, %s",
        LogTag.c_str(),
        GetNodeAttrRequest.GetNodeId(),
        GetNodeAttrRequest.GetName().Quote().c_str());

    auto request = std::make_unique<TEvService::TEvGetNodeAttrRequest>();
    request->Record = GetNodeAttrRequest;
    request->Record.SetFileSystemId(
        LeaderResponse.GetNode().GetFollowerFileSystemId());
    request->Record.SetNodeId(RootNodeId);
    request->Record.SetName(LeaderResponse.GetNode().GetFollowerNodeName());

    // forward request through tablet proxy
    ctx.Send(MakeIndexTabletProxyServiceId(), request.release());
}

////////////////////////////////////////////////////////////////////////////////

void TGetNodeAttrActor::HandleGetNodeAttrResponse(
    const TEvService::TEvGetNodeAttrResponse::TPtr& ev,
    const TActorContext& ctx)
{
    auto* msg = ev->Get();

    if (HasError(msg->GetError())) {
        HandleError(ctx, *msg->Record.MutableError());
        return;
    }

    if (LeaderResponded) {
        LOG_DEBUG(
            ctx,
            TFileStoreComponents::SERVICE,
            "GetNodeAttr succeeded in follower: %lu",
            msg->Record.GetNode().GetId());

        ReplyAndDie(ctx, std::move(msg->Record));
        return;
    }

    LOG_DEBUG(
        ctx,
        TFileStoreComponents::SERVICE,
        "GetNodeAttr succeeded in leader: %s %s",
        msg->Record.GetNode().GetFollowerFileSystemId().c_str(),
        msg->Record.GetNode().GetFollowerNodeName().Quote().c_str());

    if (!MultiTabletForwardingEnabled
            || msg->Record.GetNode().GetFollowerFileSystemId().Empty())
    {
        ReplyAndDie(ctx, std::move(msg->Record));
        return;
    }

    LeaderResponded = true;
    LeaderResponse = std::move(msg->Record);
    GetNodeAttrInFollower(ctx);
}

////////////////////////////////////////////////////////////////////////////////

void TGetNodeAttrActor::HandlePoisonPill(
    const TEvents::TEvPoisonPill::TPtr& ev,
    const TActorContext& ctx)
{
    Y_UNUSED(ev);
    HandleError(ctx, MakeError(E_REJECTED, "request cancelled"));
}

////////////////////////////////////////////////////////////////////////////////

void TGetNodeAttrActor::ReplyAndDie(
    const TActorContext& ctx,
    NProto::TGetNodeAttrResponse followerResponse)
{
    auto response = std::make_unique<TEvService::TEvGetNodeAttrResponse>();
    response->Record = std::move(followerResponse);

    NCloud::Reply(ctx, *RequestInfo, std::move(response));
    Die(ctx);
}

void TGetNodeAttrActor::HandleError(
    const TActorContext& ctx,
    NProto::TError error)
{
    auto response = std::make_unique<TEvService::TEvGetNodeAttrResponse>(
        std::move(error));
    NCloud::Reply(ctx, *RequestInfo, std::move(response));
    Die(ctx);
}

////////////////////////////////////////////////////////////////////////////////

STFUNC(TGetNodeAttrActor::StateWork)
{
    switch (ev->GetTypeRewrite()) {
        HFunc(TEvents::TEvPoisonPill, HandlePoisonPill);

        HFunc(
            TEvService::TEvGetNodeAttrResponse,
            HandleGetNodeAttrResponse);

        default:
            HandleUnexpectedEvent(ev, TFileStoreComponents::SERVICE_WORKER);
            break;
    }
}

}   // namespace

////////////////////////////////////////////////////////////////////////////////

void TStorageServiceActor::HandleGetNodeAttr(
    const TEvService::TEvGetNodeAttrRequest::TPtr& ev,
    const TActorContext& ctx)
{
    auto* msg = ev->Get();

    if (msg->Record.GetName().Empty()) {
        // handle creation by NodeId can be handled directly by the follower
        ForwardRequestToFollower<TEvService::TGetNodeAttrMethod>(
            ctx,
            ev,
            ExtractShardNo(msg->Record.GetNodeId()));
        return;
    }

    const auto& clientId = GetClientId(msg->Record);
    const auto& sessionId = GetSessionId(msg->Record);
    const ui64 seqNo = GetSessionSeqNo(msg->Record);

    auto* session = State->FindSession(sessionId, seqNo);
    if (!session || session->ClientId != clientId || !session->SessionActor) {
        auto response = std::make_unique<TEvService::TEvGetNodeAttrResponse>(
            ErrorInvalidSession(clientId, sessionId, seqNo));
        return NCloud::Reply(ctx, *ev, std::move(response));
    }

    auto [cookie, inflight] = CreateInFlightRequest(
        TRequestInfo(ev->Sender, ev->Cookie, msg->CallContext),
        session->MediaKind,
        session->RequestStats,
        ctx.Now());

    InitProfileLogRequestInfo(inflight->ProfileLogRequest, msg->Record);

    auto requestInfo = CreateRequestInfo(SelfId(), cookie, msg->CallContext);

    auto actor = std::make_unique<TGetNodeAttrActor>(
        std::move(requestInfo),
        std::move(msg->Record),
        msg->Record.GetFileSystemId(),
        session->RequestStats,
        ProfileLog,
        StorageConfig->GetMultiTabletForwardingEnabled());

    NCloud::Register(ctx, std::move(actor));
}

}   // namespace NCloud::NFileStore::NStorage
