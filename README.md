# bitpack

[![CircleCI](https://circleci.com/gh/oyi812/bitpack.svg?style=svg)](https://circleci.com/gh/oyi812/bitpack)
[![Go Report Card](https://goreportcard.com/badge/github.com/oyi812/bitpack?style=flat-square)](https://goreportcard.com/report/github.com/oyi812/bitpack)
[![Coverage](https://codecov.io/gh/oyi812/bitpack/branch/master/graph/badge.svg)](https://codecov.io/gh/oyi812/bitpack)
[![LICENSE](https://img.shields.io/badge/license-MIT-lightgrey?style=flat-square)](https://github.com/oyi812/bitpack/blob/master/LICENSE)

Efficiently encode/decode tuples of small values of fixed or prefixed sizes into/from byte slices.

## Tests

See test file for example useage.

	go test .

## Benchmarks

Benchmark against go's standard library var int binary encoding.

	go test -bench .
