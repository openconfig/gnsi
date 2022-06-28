load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

### Bazel rules for many languages to compile PROTO into gRPC libraries
http_archive(
    name = "rules_proto_grpc",
    sha256 = "507e38c8d95c7efa4f3b1c0595a8e8f139c885cb41a76cab7e20e4e67ae87731",
    strip_prefix = "rules_proto_grpc-4.1.1",
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/4.1.1.tar.gz"],
)

load(
    "@rules_proto_grpc//:repositories.bzl",
    "bazel_gazelle",
    "io_bazel_rules_go",
    "rules_proto_grpc_repos",
    "rules_proto_grpc_toolchains",
)

rules_proto_grpc_toolchains()

rules_proto_grpc_repos()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

### Golang
io_bazel_rules_go()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(go_version = "1.18")

# gazelle:repo bazel_gazelle
bazel_gazelle()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "com_github_openconfig_gnoi",
    importpath = "github.com/openconfig/gnoi",
    sum = "h1:koTAGmBf6l9XvBqinTC484NSBuCCUaz01136ofgShgo=",
    version = "v0.0.0-20220615151501-18d7d5153945",
    build_file_generation = "on",
    build_directives = [
        "gazelle:proto_import_prefix github.com/openconfig/gnoi",
    ],
)

gazelle_dependencies()

load("@rules_proto_grpc//go:repositories.bzl", rules_proto_grpc_go_repos = "go_repos")

rules_proto_grpc_go_repos()

### C++
load("@rules_proto_grpc//cpp:repositories.bzl", rules_proto_grpc_cpp_repos = "cpp_repos")

rules_proto_grpc_cpp_repos()

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

