load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "bar",
    srcs = ["bar.go"],
    importpath = "github.com/bozaro/rules_go_xdefs/bar",
    visibility = ["//visibility:public"],
    x_defs = {"Bar": "Bar"},
    deps = [
        "//baz",
        "//foo",
    ],
)
