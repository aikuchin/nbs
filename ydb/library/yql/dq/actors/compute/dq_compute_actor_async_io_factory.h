#pragma once
#include "dq_compute_actor_async_io.h"

#include <ydb/library/yql/dq/actors/protos/dq_events.pb.h>
#include <ydb/library/yql/dq/common/dq_common.h>

#include <ydb/library/yql/minikql/computation/mkql_computation_node_holders.h>
#include <ydb/library/yql/utils/yql_panic.h>

#include <type_traits>

namespace NYql::NDq {

template <class T>
concept TCastsToAsyncInputPair =
    std::is_convertible_v<T, std::pair<IDqComputeActorAsyncInput*, NActors::IActor*>>;

template <class T, class TProto>
concept TSourceCreatorFunc = requires(T f, TProto&& settings, IDqAsyncIoFactory::TSourceArguments&& args) {
    { f(std::move(settings), std::move(args)) } -> TCastsToAsyncInputPair;
};

template <class T, class TProto>
concept TSourceCreatorFuncMove = requires(T f, TProto* settings, IDqAsyncIoFactory::TSourceArguments&& args) {
    { f(settings, std::move(args)) } -> TCastsToAsyncInputPair;
};

template <class T, class TProto>
concept TInputTransformCreatorFunc = requires(T f, TProto&& settings, IDqAsyncIoFactory::TInputTransformArguments&& args) {
    { f(std::move(settings), std::move(args)) } -> TCastsToAsyncInputPair;
};

template <class T>
concept TCastsToAsyncOutputPair =
    std::is_convertible_v<T, std::pair<IDqComputeActorAsyncOutput*, NActors::IActor*>>;

template <class T, class TProto>
concept TSinkCreatorFunc = requires(T f, TProto&& settings, IDqAsyncIoFactory::TSinkArguments&& args) {
    { f(std::move(settings), std::move(args)) } -> TCastsToAsyncOutputPair;
};

template <class T, class TProto>
concept TOutputTransformCreatorFunc = requires(T f, TProto&& settings, IDqAsyncIoFactory::TOutputTransformArguments&& args) {
    { f(std::move(settings), std::move(args)) } -> TCastsToAsyncOutputPair;
};

class TDqAsyncIoFactory : public IDqAsyncIoFactory {
public:
    using TSourceCreatorFunction = std::function<std::pair<IDqComputeActorAsyncInput*, NActors::IActor*>(TSourceArguments&& args)>;
    using TSinkCreatorFunction = std::function<std::pair<IDqComputeActorAsyncOutput*, NActors::IActor*>(TSinkArguments&& args)>;
    using TInputTransformCreatorFunction = std::function<std::pair<IDqComputeActorAsyncInput*, NActors::IActor*>(TInputTransformArguments&& args)>;
    using TOutputTransformCreatorFunction = std::function<std::pair<IDqComputeActorAsyncOutput*, NActors::IActor*>(TOutputTransformArguments&& args)>;

    // Registration
    void RegisterSource(const TString& type, TSourceCreatorFunction creator);


    // This function expects to have Arena where SourceSettings are allocated or will be allocated
    template <class TProtoMsg, TSourceCreatorFuncMove<TProtoMsg> TCreatorFunc>
    void RegisterSourceWithPtr(const TString& type, TCreatorFunc creator) {  // WithArena
        RegisterSource(type,
            [creator = std::move(creator), type](TSourceArguments&& args)
            {
                YQL_ENSURE(args.Arena, "args are expected to have Arena for SourceSettings");
                if (args.SourceSettings != nullptr) {
                    const TProtoMsg* settingsPtr = dynamic_cast<const TProtoMsg*>(args.SourceSettings);
                    YQL_ENSURE(settingsPtr);
                    return creator(settingsPtr, std::move(args));
                } else {
                    const google::protobuf::Any& settingsAny = args.InputDesc.GetSource().GetSettings();
                    YQL_ENSURE(settingsAny.Is<TProtoMsg>(),
                        "Source \"" << type << "\" settings are expected to have protobuf type " << TProtoMsg::descriptor()->full_name()
                        << ", but got " << settingsAny.type_url());
                    TProtoMsg* settingsPtr = args.Arena->Allocate<TProtoMsg>();
                    YQL_ENSURE(settingsAny.UnpackTo(settingsPtr), "Failed to unpack settings of type \"" << type << "\"");
                    return creator(std::move(settingsPtr), std::move(args));
                }
        });
    }

    template <class TProtoMsg, TSourceCreatorFunc<TProtoMsg> TCreatorFunc>
    void RegisterSource(const TString& type, TCreatorFunc creator) {
        RegisterSource(type,
            [creator = std::move(creator), type](TSourceArguments&& args)
            {
                const google::protobuf::Any& settingsAny = args.InputDesc.GetSource().GetSettings();
                YQL_ENSURE(settingsAny.Is<TProtoMsg>(),
                    "Source \"" << type << "\" settings are expected to have protobuf type " << TProtoMsg::descriptor()->full_name()
                    << ", but got " << settingsAny.type_url());
                TProtoMsg settings;
                YQL_ENSURE(settingsAny.UnpackTo(&settings), "Failed to unpack settings of type \"" << type << "\"");
                return creator(std::move(settings), std::move(args));
        });
    }

    void RegisterSink(const TString& type, TSinkCreatorFunction creator);

    template <class TProtoMsg, TSinkCreatorFunc<TProtoMsg> TCreatorFunc>
    void RegisterSink(const TString& type, TCreatorFunc creator) {
        RegisterSink(type,
            [creator = std::move(creator), type](TSinkArguments&& args)
            {
                const google::protobuf::Any& settingsAny = args.OutputDesc.GetSink().GetSettings();
                YQL_ENSURE(settingsAny.Is<TProtoMsg>(),
                    "Sink \"" << type << "\" settings are expected to have protobuf type " << TProtoMsg::descriptor()->full_name()
                    << ", but got " << settingsAny.type_url());
                TProtoMsg settings;
                YQL_ENSURE(settingsAny.UnpackTo(&settings), "Failed to unpack settings of type \"" << type << "\"");
                return creator(std::move(settings), std::move(args));
        });
    }

    void RegisterInputTransform(const TString& type, TInputTransformCreatorFunction creator);

    template <class TProtoMsg, TInputTransformCreatorFunc<TProtoMsg> TCreatorFunc>
    void RegisterInputTransform(const TString& type, TCreatorFunc creator) {
        RegisterInputTransform(type,
            [creator = std::move(creator), type](TInputTransformArguments&& args)
            {
                const google::protobuf::Any& settingsAny = args.InputDesc.GetTransform().GetSettings();
                YQL_ENSURE(settingsAny.Is<TProtoMsg>(),
                    "Input transform \"" << type << "\" settings are expected to have protobuf type " << TProtoMsg::descriptor()->full_name()
                    << ", but got " << settingsAny.type_url());
                TProtoMsg settings;
                YQL_ENSURE(settingsAny.UnpackTo(&settings), "Failed to unpack settings of type \"" << type << "\"");
                return creator(std::move(settings), std::move(args));
        });
    }

    void RegisterOutputTransform(const TString& type, TOutputTransformCreatorFunction creator);

    template <class TProtoMsg, TOutputTransformCreatorFunc<TProtoMsg> TCreatorFunc>
    void RegisterOutputTransform(const TString& type, TCreatorFunc creator) {
        RegisterOutputTransform(type,
            [creator = std::move(creator), type](TOutputTransformArguments&& args)
            {
                const google::protobuf::Any& settingsAny = args.OutputDesc.GetTransform().GetSettings();
                YQL_ENSURE(settingsAny.Is<TProtoMsg>(),
                    "Output transform \"" << type << "\" settings are expected to have protobuf type " << TProtoMsg::descriptor()->full_name()
                    << ", but got " << settingsAny.type_url());
                TProtoMsg settings;
                YQL_ENSURE(settingsAny.UnpackTo(&settings), "Failed to unpack settings of type \"" << type << "\"");
                return creator(std::move(settings), std::move(args));
        });
    }

    // Creation
    std::pair<IDqComputeActorAsyncInput*, NActors::IActor*> CreateDqSource(TSourceArguments&& args) const override;
    std::pair<IDqComputeActorAsyncOutput*, NActors::IActor*> CreateDqSink(TSinkArguments&& args) const override;
    std::pair<IDqComputeActorAsyncInput*, NActors::IActor*> CreateDqInputTransform(TInputTransformArguments&& args) override;
    std::pair<IDqComputeActorAsyncOutput*, NActors::IActor*> CreateDqOutputTransform(TOutputTransformArguments&& args) override;

private:
    THashMap<TString, TSourceCreatorFunction> SourceCreatorsByType;
    THashMap<TString, TSinkCreatorFunction> SinkCreatorsByType;
    THashMap<TString, TInputTransformCreatorFunction> InputTransformCreatorsByType;
    THashMap<TString, TOutputTransformCreatorFunction> OutputTransformCreatorsByType;
};

} // namespace NYql::NDq
