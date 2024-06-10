# go-zig-cc
This repo demonstrate the usage of zig cc as cross compiler for go with usage of cgo

## CGO
> Cgo lets Go packages call C code. Given a Go source file written with some special features, cgo outputs Go and C files that can be combined into a single Go package. 

(Andrew Gerrand on [go dev blog](https://go.dev/blog/cgo))

The C files have to be compiled using a compiler like GCC or Clang.

## Zig
If you want to have an easy, quick and reliable cross-compilation to other architectures and operating systems, you can use Zig as a cross compiler.

## Go and Zig
To use Zig for cross compiling your Go code, you have to 
- set `GOOS` to your desired target OS
- set `GOARCH` to your desired target architecture
- set `CC` to the respective Zig command

As a result, cross compiling to ARM64 architecture statically linked executable on linux would result in:

`CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC="zig cc -target aarch64-linux-musl" go build`