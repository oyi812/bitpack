# bitpack

[![Go Report Card](https://goreportcard.com/badge/github.com/oyi812/bitpack?style=flat-square)](https://goreportcard.com/report/github.com/oyi812/bitpack)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/oyi812/bitpack)](https://pkg.go.dev/github.com/oyi812/bitpack)
[![LICENSE](https://img.shields.io/badge/license-MIT-lightgrey?style=flat-square)](https://github.com/oyi812/bitpack/blob/master/LICENSE)

Efficiently encode/decode int tuples of predefined or prefixed sizes into/from byte slices.

## Tests

See test file for example useage.

	go test .

## Benchmarks

Benchmark against go's standard library var int binary encoding.

	go test -bench .

## Build

      go get -v -t ./...
      go test -v ./...

