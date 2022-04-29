load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "qux",
    srcs = ["qux.go"],
    importpath = "github.com/bozaro/rules_go_xdefs/qux",
    visibility = ["//visibility:public"],
    x_defs = {"Qux": "Qux"},
)
