#pragma once

#include "last_error.h"
#include "service_tickets.h"
#include "settings.h"
#include "roles/roles.h"

#include <library/cpp/tvmauth/client/client_status.h>
#include <library/cpp/tvmauth/client/logger.h>

#include <library/cpp/tvmauth/deprecated/service_context.h>
#include <library/cpp/tvmauth/deprecated/user_context.h>
#include <library/cpp/tvmauth/src/utils.h>

#include <util/datetime/base.h>
#include <util/generic/hash.h>
#include <util/generic/maybe.h>
#include <util/generic/noncopyable.h>
#include <util/generic/ptr.h>

#include <array>
#include <atomic>

namespace NTvmAuth {

    class TAllUserContexts: public TAtomicRefCount<TAllUserContexts> {
    public:
        TAllUserContexts(TStringBuf publicKeys);

        TUserContextPtr Get(EBlackboxEnv env) const;

    private:
        std::array<TUserContextPtr, 5> Ctx_;
    };
    using TAllUserContextsPtr = TIntrusiveConstPtr<TAllUserContexts>;

    class TAsyncUpdaterBase: public TAtomicRefCount<TAsyncUpdaterBase>, protected TLastError, TNonCopyable {
    public:
        TAsyncUpdaterBase();
        virtual ~TAsyncUpdaterBase() = default;

        virtual TClientStatus GetStatus() const = 0;
        virtual TString GetServiceTicketFor(const TClientSettings::TAlias& dst) const;
        virtual TString GetServiceTicketFor(const TTvmId dst) const;
        virtual TCheckedServiceTicket CheckServiceTicket(TStringBuf ticket, const TServiceContext::TCheckFlags& flags = TServiceContext::TCheckFlags{}) const;
        virtual TCheckedUserTicket CheckUserTicket(TStringBuf ticket, TMaybe<EBlackboxEnv> overrideEnv = {}) const;
        virtual NRoles::TRolesPtr GetRoles() const;

        TServiceTicketsPtr GetCachedServiceTickets() const;
        TServiceContextPtr GetCachedServiceContext() const;
        TUserContextPtr GetCachedUserContext(TMaybe<EBlackboxEnv> overridenEnv = {}) const;

        TInstant GetUpdateTimeOfPublicKeys() const;
        TInstant GetUpdateTimeOfServiceTickets() const;
        TInstant GetUpdateTimeOfRoles() const;
        TInstant GetInvalidationTimeOfPublicKeys() const;
        TInstant GetInvalidationTimeOfServiceTickets() const;

        bool ArePublicKeysInvalid(TInstant now) const;
        bool AreServiceTicketsInvalid(TInstant now) const;
        static bool IsInvalid(TInstant invTime, TInstant now);

    protected:
        void SetBbEnv(EBlackboxEnv original, TMaybe<EBlackboxEnv> overrided = {});

        void SetServiceTickets(TServiceTicketsPtr c);
        void SetServiceContext(TServiceContextPtr c);
        void SetUserContext(TStringBuf publicKeys);
        void SetUpdateTimeOfPublicKeys(TInstant ins);
        void SetUpdateTimeOfServiceTickets(TInstant ins);
        void SetUpdateTimeOfRoles(TInstant ins);

        void SetInited(bool value);
        bool IsInited() const;

        static bool IsServiceTicketMapOk(TServiceTicketsPtr c, size_t expectedTicketCount, bool strict);

    protected:
        struct TPairTicketsErrors {
            TServiceTickets::TMapIdStr Tickets;
            TServiceTickets::TMapIdStr Errors;

            bool operator==(const TPairTicketsErrors& o) const {
                return Tickets == o.Tickets && Errors == o.Errors;
            }
        };

        struct TStateDurations {
            TDuration RefreshPeriod;
            TDuration Expiring;
            TDuration Invalid;
        };

        TStateDurations ServiceTicketsDurations_;
        TStateDurations PublicKeysDurations_;

    protected:
        virtual void StartTvmClientStopping() const {
        }
        virtual bool IsTvmClientStopped() const {
            return true;
        }
        friend class NTvmAuth::NInternal::TClientCaningKnife;

    private:
        struct TEnvs {
            TMaybe<EBlackboxEnv> Original;
            TMaybe<EBlackboxEnv> Overrided;
        };
        static_assert(sizeof(TEnvs) <= 8, "Small struct is easy to store as atomic");
        std::atomic<TEnvs> Envs_ = {{}};

        NUtils::TProtectedValue<TServiceTicketsPtr> ServiceTickets_;
        NUtils::TProtectedValue<TServiceContextPtr> ServiceContext_;
        NUtils::TProtectedValue<TAllUserContextsPtr> AllUserContexts_;
        NUtils::TProtectedValue<TInstant> PublicKeysTime_;
        NUtils::TProtectedValue<TInstant> ServiceTicketsTime_;
        NUtils::TProtectedValue<TInstant> RolesTime_;

        std::atomic_bool Inited_{false};
    };
    using TAsyncUpdaterPtr = TIntrusiveConstPtr<TAsyncUpdaterBase>;
}
