#include <cloud/storage/core/libs/diagnostics/cgroup_stats_fetcher.h>
#include <cloud/storage/core/libs/diagnostics/logging.h>
#include <cloud/storage/core/libs/diagnostics/monitoring.h>

#include <library/cpp/getopt/small/last_getopt.h>
#include <library/cpp/monlib/dynamic_counters/counters.h>

namespace {

const TString DefaultPath = "/sys/fs/cgroup/cpu/system.slice/nbs.service/cpuacct.wait";
constexpr ui32 DefaultPollPeriod = 1;
constexpr ui32 DefaultNumSamples = 0;
const TString DefaultComponentName = "STORAGE_CGROUP";

////////////////////////////////////////////////////////////////////////////////

struct TOptions
{
    TString Path = DefaultPath;
    ui32 PollPeriod = DefaultPollPeriod;
    ui32 SampleCount = DefaultNumSamples;
    TString ComponentName = DefaultComponentName;

    TOptions(int argc, const char** argv)
    {
        using namespace NLastGetopt;

        TOpts opts;
        opts.AddHelpOption();

        opts.AddLongOption("file", "CPU wait stats file location")
            .RequiredArgument("STRING")
            .StoreResult(&Path);

        opts.AddLongOption("poll-period", "metric poll period in seconds")
            .RequiredArgument("INTEGER")
            .StoreResult(&PollPeriod);

        opts.AddLongOption("samples", "number of samples to take")
            .RequiredArgument("INTEGER")
            .StoreResult(&SampleCount);

        opts.AddLongOption("component", "component name")
            .RequiredArgument("STRING")
            .StoreResult(&ComponentName);

        TOptsParseResultException(&opts, argc, argv);
    }
};

}   // namespace

////////////////////////////////////////////////////////////////////////////////

int main(int argc, const char** argv)
{
    TOptions options(argc, argv);

    auto logging =
        NCloud::CreateLoggingService("console", NCloud::TLogSettings{});
    auto Log = logging->CreateLog(options.ComponentName);
    auto monitoring = NCloud::CreateMonitoringServiceStub();
    auto statsFetcher = NCloud::NStorage::CreateCgroupStatsFetcher(
        options.ComponentName,
        logging,
        monitoring,
        options.Path,
        {
            .CountersGroupName = options.ComponentName,
            .ComponentGroupName = "server",
            .CounterName = "CpuWaitFailure",
        });

    statsFetcher->Start();

    TDuration pollInterval = TDuration::Seconds(options.PollPeriod);

    auto numSamples = options.SampleCount;

    while (!options.SampleCount || numSamples--) {
        Sleep(pollInterval);

        auto waitTime = 100 * statsFetcher->GetCpuWait().MicroSeconds();
        auto interval = pollInterval.MicroSeconds();

        Cout << (waitTime / interval) << Endl;
    }

    return 0;
}
