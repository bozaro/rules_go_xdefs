load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "foo",
    srcs = ["foo.go"],
    importpath = "github.com/bozaro/rules_go_xdefs/foo",
    visibility = ["//visibility:public"],
)

go_test(
    name = "foo_test",
    srcs = ["foo_test.go"],
    embed = [":foo"],
    importpath = "github.com/bozaro/rules_go_xdefs/foo",
    pure = "on",
    deps = ["//bar"],
)
