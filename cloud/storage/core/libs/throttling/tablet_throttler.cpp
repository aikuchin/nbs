#include "tablet_throttler.h"

#include "tablet_throttler_logger.h"
#include "tablet_throttler_policy.h"

#include <cloud/storage/core/libs/common/context.h>
#include <cloud/storage/core/libs/common/error.h>

#include <library/cpp/actors/core/actor.h>
#include <library/cpp/actors/core/events.h>

#include <util/datetime/base.h>
#include <util/generic/list.h>

namespace NCloud {

////////////////////////////////////////////////////////////////////////////////

constexpr TDuration MinPostponeQueueFlushInterval = TDuration::MilliSeconds(1);

////////////////////////////////////////////////////////////////////////////////

class TTabletThrottler final
    : public ITabletThrottler
{
private:
    NActors::IActor& Owner;
    ITabletThrottlerLogger& Logger;
    ITabletThrottlerPolicy& Policy;

    struct TPostponedRequest
    {
        TThrottlingRequestInfo Info;
        TCallContextBasePtr CallContext;
        NActors::IEventHandlePtr Event;
    };
    // TODO: replace with a ring buffer
    TList<TPostponedRequest> PostponedRequests;
    bool PostponedQueueFlushScheduled = false;
    bool PostponedQueueFlushInProgress = false;

public:
    TTabletThrottler(
            NActors::IActor& owner,
            ITabletThrottlerLogger& logger,
            ITabletThrottlerPolicy& policy)
        : Owner(owner)
        , Logger(logger)
        , Policy(policy)
    {}

    ui64 GetPostponedRequestsCount() const override
    {
        return PostponedRequests.size();
    }

    void ResetPolicy(ITabletThrottlerPolicy& policy) override
    {
        Policy = policy;
    }

    void OnShutDown(const NActors::TActorContext& ctx) override
    {
        PostponedQueueFlushScheduled = false;

        while (PostponedRequests.size()) {
            auto& x = PostponedRequests.front();
            Advance(ctx, x);

            TAutoPtr<NActors::IEventHandle> ev = x.Event.release();
            Owner.Receive(ev);

            Y_ABORT_UNLESS(!PostponedQueueFlushScheduled);
            PostponedRequests.pop_front();
        }
    }

    void StartFlushing(const NActors::TActorContext& ctx) override
    {
        Y_DEBUG_ABORT_UNLESS(PostponedQueueFlushScheduled);
        PostponedQueueFlushScheduled = false;
        PostponedQueueFlushInProgress = true;

        while (PostponedRequests.size()) {
            auto& x = PostponedRequests.front();
            Policy.OnPostponedEvent(ctx.Now(), x.Info);
            Advance(ctx, x);

            TAutoPtr<NActors::IEventHandle> ev = x.Event.release();
            Owner.Receive(ev);

            if (PostponedQueueFlushScheduled) {
                Y_ABORT_UNLESS(x.Event);
                break;
            }

            PostponedRequests.pop_front();
        }

        PostponedQueueFlushInProgress = false;
    }

    ETabletThrottlerStatus Throttle(
        const NActors::TActorContext& ctx,
        TCallContextBasePtr callContext,
        const TThrottlingRequestInfo& requestInfo,
        const std::function<NActors::IEventHandlePtr(void)>& eventReleaser,
        const char* methodName) override
    {
        bool rejected = false;
        if (PostponedRequests && !PostponedQueueFlushInProgress) {
            Y_DEBUG_ABORT_UNLESS(PostponedQueueFlushScheduled);

            if (Policy.TryPostpone(ctx.Now(), requestInfo)) {
                Logger.LogRequestPostponedAfterSchedule(
                    ctx,
                    *callContext,
                    GetPostponedRequestsCount(),
                    methodName);

                Postpone(
                    ctx,
                    requestInfo,
                    std::move(callContext),
                    eventReleaser());

                return ETabletThrottlerStatus::POSTPONED;
            }

            rejected = true;
        } else {
            ui64 postponeTs = callContext->GetPostponeCycles();
            TDuration queueTime;
            if (postponeTs) {
                queueTime = ctx.Now() - TInstant::MicroSeconds(postponeTs);
            }
            const auto delay =
                Policy.SuggestDelay(ctx.Now(), queueTime, requestInfo);

            if (delay.Defined()) {
                if (delay->GetValue()) {
                    Logger.LogRequestPostponedBeforeSchedule(
                        ctx,
                        *callContext,
                        *delay,
                        methodName);

                    Postpone(
                        ctx,
                        requestInfo,
                        std::move(callContext),
                        eventReleaser());

                    Y_DEBUG_ABORT_UNLESS(!PostponedQueueFlushScheduled);
                    PostponedQueueFlushScheduled = true;

                    ctx.Schedule(
                        Max(*delay, MinPostponeQueueFlushInterval),
                        new NActors::TEvents::TEvWakeup());

                    return ETabletThrottlerStatus::POSTPONED;
                } else if (PostponedQueueFlushInProgress) {
                    Logger.LogRequestAdvanced(ctx, *callContext, methodName);
                }
            } else {
                rejected = true;
            }
        }

        if (rejected) {
            return ETabletThrottlerStatus::REJECTED;
        }

        if (!PostponedQueueFlushInProgress) {
            // throttling caused no delay for this request
            Logger.UpdateDelayCounter(requestInfo.OpType, TDuration::Zero());
        }

        return ETabletThrottlerStatus::ADVANCED;
    }

private:
    void Postpone(
        const NActors::TActorContext& ctx,
        TThrottlingRequestInfo requestInfo,
        TCallContextBasePtr callContext,
        NActors::IEventHandlePtr ev)
    {
        const auto ts = ctx.Now().MicroSeconds();
        callContext->Postpone(ts);

        if (PostponedQueueFlushInProgress) {
            Y_DEBUG_ABORT_UNLESS(!PostponedRequests.front().Event);
            auto& pr = PostponedRequests.front();
            pr.Event = std::move(ev);
            pr.Info = requestInfo;
        } else {
            callContext->SetPostponeCycles(ts);
            PostponedRequests.push_back({
                requestInfo,
                std::move(callContext),
                std::move(ev)});
       }
    }

    void Advance(
        const NActors::TActorContext& ctx,
        TPostponedRequest& req)
    {
        const auto delay = req.CallContext->Advance(ctx.Now().MicroSeconds());
        Logger.LogPostponedRequestAdvanced(
            *req.CallContext,
            req.Info.OpType,
            delay);
    }

};

////////////////////////////////////////////////////////////////////////////////

ITabletThrottlerPtr CreateTabletThrottler(
    NActors::IActor& owner,
    ITabletThrottlerLogger& logger,
    ITabletThrottlerPolicy& policy)
{
    return std::make_unique<TTabletThrottler>(owner, logger, policy);
}

}   // namespace NCloud
