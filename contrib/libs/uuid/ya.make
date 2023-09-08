# Generated by devtools/yamaker from nixpkgs 22.11.

LIBRARY()

LICENSE(
    BSD-3-Clause AND
    Public-Domain
)

LICENSE_TEXTS(.yandex_meta/licenses.list.txt)

VERSION(2.39.1)

ORIGINAL_SOURCE(https://github.com/util-linux/util-linux/archive/v2.39.1.tar.gz)

PEERDIR(
    contrib/libs/libc_compat
)

ADDINCL(
    contrib/libs/uuid
    contrib/libs/uuid/include
    contrib/libs/uuid/libuuid/src
)

NO_COMPILER_WARNINGS()

NO_RUNTIME()

CFLAGS(
    -DHAVE_CONFIG_H
    -DLOCALEDIR=\"/tmp/yamaker/uuid/out/share/locale\"
    -D_PATH_RUNSTATEDIR=\"/run\"
    -D_PATH_SYSCONFSTATICDIR=\"/0sra2y18lr3h6j58qjm0w46yv36h1wjmilb09n8aimdpivdymscx/lib\"
)

SRCS(
    lib/md5.c
    lib/randutils.c
    lib/sha1.c
    libuuid/src/clear.c
    libuuid/src/compare.c
    libuuid/src/copy.c
    libuuid/src/gen_uuid.c
    libuuid/src/isnull.c
    libuuid/src/pack.c
    libuuid/src/parse.c
    libuuid/src/predefined.c
    libuuid/src/unpack.c
    libuuid/src/unparse.c
    libuuid/src/uuid_time.c
)

END()
