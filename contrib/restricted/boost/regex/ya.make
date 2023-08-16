# Generated by devtools/yamaker from nixpkgs 22.11.

LIBRARY()

LICENSE(BSL-1.0)

LICENSE_TEXTS(.yandex_meta/licenses.list.txt)

VERSION(1.83.0)

ORIGINAL_SOURCE(https://github.com/boostorg/regex/archive/boost-1.83.0.tar.gz)

PEERDIR(
    contrib/libs/icu
    contrib/restricted/boost/assert
    contrib/restricted/boost/concept_check
    contrib/restricted/boost/config
    contrib/restricted/boost/container_hash
    contrib/restricted/boost/core
    contrib/restricted/boost/integer
    contrib/restricted/boost/mpl
    contrib/restricted/boost/predef
    contrib/restricted/boost/smart_ptr
    contrib/restricted/boost/static_assert
    contrib/restricted/boost/throw_exception
    contrib/restricted/boost/type_traits
)

ADDINCL(
    GLOBAL contrib/restricted/boost/regex/include
)

NO_COMPILER_WARNINGS()

NO_UTIL()

CFLAGS(
    -DBOOST_HAS_ICU
    -DBOOST_NO_CXX98_BINDERS
)

IF (DYNAMIC_BOOST)
    CFLAGS(
        GLOBAL -DBOOST_REGEX_DYN_LINK
    )
ENDIF()

IF (OS_WINDOWS)
    CFLAGS(
        GLOBAL -DBOOST_REGEX_USE_CPP_LOCALE
    )
ENDIF()

SRCS(
    src/posix_api.cpp
    src/regex.cpp
    src/regex_debug.cpp
    src/static_mutex.cpp
    src/wide_posix_api.cpp
)

END()
