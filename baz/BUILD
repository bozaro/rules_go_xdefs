load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "baz",
    srcs = ["baz.go"],
    importpath = "github.com/bozaro/rules_go_xdefs/baz",
    visibility = ["//visibility:public"],
    x_defs = {"Baz": "Baz"},
    deps = ["//qux"],
)
