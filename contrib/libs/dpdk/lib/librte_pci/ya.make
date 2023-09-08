# Generated by devtools/yamaker.

LIBRARY()

LICENSE(BSD-3-Clause)

LICENSE_TEXTS(.yandex_meta/licenses.list.txt)

PEERDIR(
    contrib/libs/dpdk/lib/librte_eal
    contrib/libs/dpdk/lib/librte_kvargs
    contrib/libs/dpdk/lib/librte_telemetry
)

ADDINCL(
    GLOBAL contrib/libs/dpdk/config
    GLOBAL contrib/libs/dpdk/lib/librte_pci
    contrib/libs/dpdk/lib/librte_eal
    contrib/libs/dpdk/lib/librte_eal/include
    contrib/libs/dpdk/lib/librte_eal/linux/include
)

NO_COMPILER_WARNINGS()

NO_RUNTIME()

CFLAGS(
    -DALLOW_EXPERIMENTAL_API
    -DALLOW_INTERNAL_API
)

SRCS(
    rte_pci.c
)

IF (ARCH_ARM64)
    ADDINCL(contrib/libs/dpdk/lib/librte_eal/arm/include)
ELSEIF (ARCH_X86_64)
    ADDINCL(contrib/libs/dpdk/lib/librte_eal/x86/include)
ENDIF()

END()
