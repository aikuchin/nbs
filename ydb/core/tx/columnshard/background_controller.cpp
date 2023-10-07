#include "background_controller.h"
#include <ydb/core/tx/columnshard/engines/column_engine_logs.h>

namespace NKikimr::NColumnShard {

void TBackgroundController::StartTtl(const NOlap::TColumnEngineChanges& changes) {
    Y_VERIFY(TtlPortions.empty());
    TtlPortions = changes.GetTouchedPortions();
}

bool TBackgroundController::StartCompaction(const NOlap::TPlanCompactionInfo& info, const NOlap::TColumnEngineChanges& changes) {
    Y_VERIFY(ActiveCompactionInfo.emplace(info.GetPathId(), info).second);
    Y_VERIFY(CompactionInfoPortions.emplace(info.GetPathId(), changes.GetTouchedPortions()).second);
    return true;
}

THashSet<NOlap::TPortionAddress> TBackgroundController::GetConflictTTLPortions() const {
    THashSet<NOlap::TPortionAddress> result = TtlPortions;
    for (auto&& i : CompactionInfoPortions) {
        for (auto&& g : i.second) {
            Y_VERIFY(result.emplace(g).second);
        }
    }
    return result;
}

THashSet<NOlap::TPortionAddress> TBackgroundController::GetConflictCompactionPortions() const {
    THashSet<NOlap::TPortionAddress> result = TtlPortions;
    for (auto&& i : CompactionInfoPortions) {
        for (auto&& g : i.second) {
            Y_VERIFY(result.emplace(g).second);
        }
    }
    return result;
}

void TBackgroundController::CheckDeadlines() {
    for (auto&& i : ActiveCompactionInfo) {
        if (TMonotonic::Now() - i.second.GetStartTime() > NOlap::TCompactionLimits::CompactionTimeout) {
            AFL_EMERG(NKikimrServices::TX_COLUMNSHARD)("event", "deadline_compaction");
            Y_VERIFY_DEBUG(false);
        }
    }
}

void TBackgroundController::CheckDeadlinesIndexation() {
    for (auto&& i : ActiveIndexationTasks) {
        if (TMonotonic::Now() - i.second > NOlap::TCompactionLimits::CompactionTimeout) {
            AFL_EMERG(NKikimrServices::TX_COLUMNSHARD)("event", "deadline_compaction")("task_id", i.first);
            Y_VERIFY_DEBUG(false);
        }
    }
}

void TBackgroundController::StartIndexing(const NOlap::TColumnEngineChanges& changes) {
    Y_VERIFY(ActiveIndexationTasks.emplace(changes.GetTaskIdentifier(), TMonotonic::Now()).second);
}

void TBackgroundController::FinishIndexing(const NOlap::TColumnEngineChanges& changes) {
    Y_VERIFY(ActiveIndexationTasks.erase(changes.GetTaskIdentifier()));
}

TString TBackgroundController::DebugStringIndexation() const {
    TStringBuilder sb;
    sb << "{";
    sb << "task_ids=";
    for (auto&& i : ActiveIndexationTasks) {
        sb << i.first << ",";
    }
    sb << ";";
    sb << "}";
    return sb;
}

TString TBackgroundActivity::DebugString() const {
    return TStringBuilder()
        << "indexation:" << HasIndexation() << ";"
        << "compaction:" << HasCompaction() << ";"
        << "cleanup:" << HasCleanup() << ";"
        << "ttl:" << HasTtl() << ";"
        ;
}

}
