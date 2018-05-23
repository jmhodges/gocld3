workspace(name = "com_github_jmhodges_gocld3")

# required by gazelle
skylib_version = "f9b0ff1dd3d119d19b9cacbbc425a9e61759f1f5"
http_archive(
    name = "bazel_skylib",
    sha256 = "ce27a2007deda8a1de65df9de3d4cd93a5360ead43c5ff3017ae6b3a2abe485e",
    strip_prefix= "bazel-skylib-{v}".format(v=skylib_version),
    urls = [
        "https://github.com/bazelbuild/bazel-skylib/archive/{v}.tar.gz".format(v=skylib_version)
    ]
)

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "ca213b3006c8eed6b3f1ea649cab36b817901b46",
)

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.12.0/bazel-gazelle-0.12.0.tar.gz",
    sha256 = "ddedc7aaeb61f2654d7d7d4fd7940052ea992ccdb031b8f9797ed143ac7e8d43",
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()
go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@bazel_gazelle//:def.bzl", "go_repository")
gazelle_dependencies()

git_repository(
    name = "com_github_jmhodges_cld3",
    remote = "https://github.com/jmhodges/cld3",
    commit = "d5bf85496cd4cfc8348b3233da28f91195385111",
)
