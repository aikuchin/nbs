# Generated by devtools/yamaker from nixpkgs 22.11.

LIBRARY()

LICENSE(
    Apache-2.0 AND
    BSL-1.0 AND
    FSFAP
)

LICENSE_TEXTS(../LICENSE.txt)

VERSION(1.11.3)

ORIGINAL_SOURCE(https://github.com/apache/avro/archive/release-1.11.3.tar.gz)

PEERDIR(
    contrib/libs/snappy
    contrib/restricted/boost/any
    contrib/restricted/boost/asio
    contrib/restricted/boost/crc
    contrib/restricted/boost/format
    contrib/restricted/boost/iostreams
    contrib/restricted/boost/math
)

ADDINCL(
    contrib/libs/apache/avro/api
)

NO_COMPILER_WARNINGS()

NO_UTIL()

CFLAGS(
    -DAVRO_SOURCE
    -DAVRO_VERSION=\"1.11.3\"
    -DSNAPPY_CODEC_AVAILABLE
)

SRCS(
    impl/BinaryDecoder.cc
    impl/BinaryEncoder.cc
    impl/Compiler.cc
    impl/CustomAttributes.cc
    impl/DataFile.cc
    impl/FileStream.cc
    impl/Generic.cc
    impl/GenericDatum.cc
    impl/LogicalType.cc
    impl/Node.cc
    impl/NodeImpl.cc
    impl/Resolver.cc
    impl/ResolverSchema.cc
    impl/Schema.cc
    impl/Stream.cc
    impl/Types.cc
    impl/ValidSchema.cc
    impl/Validator.cc
    impl/Zigzag.cc
    impl/json/JsonDom.cc
    impl/json/JsonIO.cc
    impl/parsing/JsonCodec.cc
    impl/parsing/ResolvingDecoder.cc
    impl/parsing/Symbol.cc
    impl/parsing/ValidatingCodec.cc
)

END()
