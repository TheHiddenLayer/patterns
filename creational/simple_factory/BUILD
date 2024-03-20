load("//go:go.bzl", "go_library", "go_binary")
load("//cpp:cpp.bzl", "cpp_library", "cpp_binary")

go_library(
    name = "simple_factory_go",
    srcs = ["go/simple_factory.go"],
    importpath = "patterns/creational/simple-factory/go",
)

go_binary(
    name = "simple_factory_go_bin",
    embed = [":simple_factory_go"],
    srcs = ["go/main.go"],
)

cpp_library(
    name = "simple_factory_cpp",
    srcs = ["cpp/simple_factory.cpp"],
    hdrs = ["cpp/simple_factory.h"],
)

cpp_binary(
    name = "simple_factory_cpp_bin",
    srcs = ["cpp/main.cpp"],
    deps = [":simple_factory_cpp"],
)