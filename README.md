# Demo Repo Reproducing a Bug in io_bazel_rules_docker

This repo is a minimal example demonstrating a bug in io_bazel_rules_docker that
prevents go_image from targeting arm64 on an M1 MacBook.

If you run the echo_image, you get the warning stating that the requested image's
platform is amd64 instead of the expected arm64:

```
$ bazel run //go/echo:echo_image
INFO: Analyzed target //go/echo:echo_image (27 packages loaded, 851 targets configured).
INFO: Found 1 target...
Target //go/echo:echo_image up-to-date:
  bazel-out/darwin_arm64-fastbuild-ST-4a519fd6d3e4/bin/go/echo/echo_image-layer.tar
INFO: Elapsed time: 1.550s, Critical Path: 0.11s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
Loaded image ID: sha256:3650e8533d18cdecc09475b7de489dc60213bb970fb97de3d9086308bd6311ba
Tagging 3650e8533d18cdecc09475b7de489dc60213bb970fb97de3d9086308bd6311ba as bazel/go/echo:echo_image
WARNING: The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested
```

The fix is to add an architecture flag which was done in this [PR](https://github.com/bazelbuild/rules_docker/pull/2150).

# Usage

```
$ bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_arm64 //go/echo:echo_image
```
