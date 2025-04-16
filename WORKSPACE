workspace(name = "com_github_openconfig_gnsi")

load("//:gnsi_deps.bzl", "gnsi_deps")

gnsi_deps()

load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    cc = True,
    go = True,
    grpc = True,
)

load("@bazel_features//:deps.bzl", "bazel_features_deps")

bazel_features_deps()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("//:gnsi_go_deps.bzl", "gnsi_go_deps")

# gazelle:repository_macro gnsi_go_deps.bzl%gnsi_go_deps
gnsi_go_deps()

go_rules_dependencies()

go_register_toolchains(version = "1.24.1")

gazelle_dependencies()

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

# Required by grpc
load("@build_bazel_apple_support//lib:repositories.bzl", "apple_support_dependencies")
load("@build_bazel_rules_apple//apple:repositories.bzl", "apple_rules_dependencies")
load("@rules_python//python:repositories.bzl", "py_repositories")

py_repositories()

apple_rules_dependencies(ignore_version_differences = False)

apple_support_dependencies()
