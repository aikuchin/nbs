#include <ydb/core/tx/schemeshard/ut_helpers/helpers.h>
#include <ydb/core/tx/schemeshard/schemeshard_billing_helpers.h>
#include <ydb/core/testlib/tablet_helpers.h>

#include <ydb/core/tx/datashard/datashard.h>
#include <ydb/core/metering/metering.h>

using namespace NKikimr;
using namespace NSchemeShard;
using namespace NSchemeShardUT_Private;

Y_UNIT_TEST_SUITE(IndexBuildTest) {
    Y_UNIT_TEST(BaseCase) {
        TTestBasicRuntime runtime;
        TTestEnv env(runtime);
        ui64 txId = 100;

        TestCreateExtSubDomain(runtime, ++txId,  "/MyRoot",
                               "Name: \"ResourceDB\"");
        env.TestWaitNotification(runtime, txId);

        TestAlterExtSubDomain(runtime, ++txId,  "/MyRoot",
                              "StoragePools { "
                              "  Name: \"pool-1\" "
                              "  Kind: \"pool-kind-1\" "
                              "} "
                              "StoragePools { "
                              "  Name: \"pool-2\" "
                              "  Kind: \"pool-kind-2\" "
                              "} "
                              "PlanResolution: 50 "
                              "Coordinators: 1 "
                              "Mediators: 1 "
                              "TimeCastBucketsPerMediator: 2 "
                              "ExternalSchemeShard: true "
                              "Name: \"ResourceDB\"");
        env.TestWaitNotification(runtime, txId);

        const auto attrs = AlterUserAttrs({
            {"cloud_id", "CLOUD_ID_VAL"},
            {"folder_id", "FOLDER_ID_VAL"},
            {"database_id", "DATABASE_ID_VAL"}
        });

        TestCreateExtSubDomain(runtime, ++txId, "/MyRoot", Sprintf(R"(
            Name: "ServerLessDB"
            ResourcesDomainKey {
                SchemeShard: %lu
                PathId: 2
            }
        )", TTestTxConfig::SchemeShard), attrs);
        env.TestWaitNotification(runtime, txId);

        TString alterData = TStringBuilder()
            << "PlanResolution: 50 "
            << "Coordinators: 1 "
            << "Mediators: 1 "
            << "TimeCastBucketsPerMediator: 2 "
            << "ExternalSchemeShard: true "
            << "ExternalHive: false "
            << "Name: \"ServerLessDB\" "
            << "StoragePools { "
            << "  Name: \"pool-1\" "
            << "  Kind: \"pool-kind-1\" "
            << "} ";
        TestAlterExtSubDomain(runtime, ++txId,  "/MyRoot", alterData);
        env.TestWaitNotification(runtime, txId);

        ui64 tenantSchemeShard = 0;
        TestDescribeResult(DescribePath(runtime, "/MyRoot/ServerLessDB"),
                           {NLs::PathExist,
                            NLs::IsExternalSubDomain("ServerLessDB"),
                            NLs::ExtractTenantSchemeshard(&tenantSchemeShard)});

        // Just create main table
        TestCreateTable(runtime, tenantSchemeShard, ++txId, "/MyRoot/ServerLessDB", R"(
              Name: "Table"
              Columns { Name: "key"     Type: "Uint32" }
              Columns { Name: "index"   Type: "Uint32" }
              Columns { Name: "value"   Type: "Utf8"   }
              KeyColumnNames: ["key"]
        )");
        env.TestWaitNotification(runtime, txId, tenantSchemeShard);

        auto fnWriteRow = [&] (ui64 tabletId, ui32 key, ui32 index, TString value, const char* table) {
            TString writeQuery = Sprintf(R"(
                (
                    (let key   '( '('key   (Uint32 '%u ) ) ) )
                    (let row   '( '('index (Uint32 '%u ) )  '('value (Utf8 '%s) ) ) )
                    (return (AsList (UpdateRow '__user__%s key row) ))
                )
            )", key, index, value.c_str(), table);
            NKikimrMiniKQL::TResult result;
            TString err;
            NKikimrProto::EReplyStatus status = LocalMiniKQL(runtime, tabletId, writeQuery, result, err);
            UNIT_ASSERT_VALUES_EQUAL(err, "");
            UNIT_ASSERT_VALUES_EQUAL(status, NKikimrProto::EReplyStatus::OK);;
        };
        for (ui32 delta = 0; delta < 101; ++delta) {
            fnWriteRow(TTestTxConfig::FakeHiveTablets + 6, 1 + delta, 1000 + delta, "aaaa", "Table");
        }

        runtime.SetLogPriority(NKikimrServices::TX_DATASHARD, NLog::PRI_TRACE);
        runtime.SetLogPriority(NKikimrServices::BUILD_INDEX, NLog::PRI_TRACE);

        TestDescribeResult(DescribePath(runtime, tenantSchemeShard, "/MyRoot/ServerLessDB/Table"),
                           {NLs::PathExist,
                            NLs::IndexesCount(0),
                            NLs::PathVersionEqual(3)});

        TStringBuilder meteringMessages;
        auto grabMeteringMessage = [&meteringMessages](TTestActorRuntimeBase&, TAutoPtr<IEventHandle> &ev) -> auto {
            if (ev->Type == NMetering::TEvMetering::TEvWriteMeteringJson::EventType) {
                auto *msg = ev->Get<NMetering::TEvMetering::TEvWriteMeteringJson>();
                Cerr << "grabMeteringMessage has happened" << Endl;
                meteringMessages << msg->MeteringJson;
            }

            return TTestActorRuntime::EEventAction::PROCESS;
        };

        runtime.SetObserverFunc(grabMeteringMessage);

        Ydb::TypedValue defaultValue;
        defaultValue.mutable_type()->set_type_id(Ydb::Type::UINT64);
        defaultValue.mutable_value()->set_uint64_value(10);

        TestBuildColumn(runtime, ++txId, tenantSchemeShard, "/MyRoot/ServerLessDB", "/MyRoot/ServerLessDB/Table", "DefaultValue", defaultValue, Ydb::StatusIds::SUCCESS);
        // ui64 buildIndexId = txId;

        auto listing = TestListBuildIndex(runtime, tenantSchemeShard, "/MyRoot/ServerLessDB");
        Y_ASSERT(listing.EntriesSize() == 1);

        env.TestWaitNotification(runtime, txId, tenantSchemeShard);

        auto descr = TestGetBuildIndex(runtime, tenantSchemeShard, "/MyRoot/ServerLessDB", txId);
        Y_ASSERT(descr.GetIndexBuild().GetState() == Ydb::Table::IndexBuildState::STATE_DONE);
/*
        const TString meteringData = R"({"usage":{"start":0,"quantity":179,"finish":0,"unit":"request_unit","type":"delta"},"tags":{},"id":"106-9437197-2-101-1818-101-1818","cloud_id":"CLOUD_ID_VAL","source_wt":0,"source_id":"sless-docapi-ydb-ss","resource_id":"DATABASE_ID_VAL","schema":"ydb.serverless.requests.v1","folder_id":"FOLDER_ID_VAL","version":"1.0.0"})";

        UNIT_ASSERT_NO_DIFF(meteringMessages, meteringData + "\n");

        TestDescribeResult(DescribePath(runtime, tenantSchemeShard, "/MyRoot/ServerLessDB/Table"),
                           {NLs::PathExist,
                            NLs::IndexesCount(1),
                            NLs::PathVersionEqual(6)});

        TestDescribeResult(DescribePath(runtime, tenantSchemeShard, "/MyRoot/ServerLessDB/Table/index1", true, true, true),
                           {NLs::PathExist,
                            NLs::IndexState(NKikimrSchemeOp::EIndexState::EIndexStateReady)});

        TestForgetBuildIndex(runtime, ++txId, tenantSchemeShard, "/MyRoot/ServerLessDB", buildIndexId);
        listing = TestListBuildIndex(runtime, tenantSchemeShard, "/MyRoot/ServerLessDB");
        Y_ASSERT(listing.EntriesSize() == 0);
*/
    }
}
