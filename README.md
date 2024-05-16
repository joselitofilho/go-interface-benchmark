# Go Interface Benchmark

[![Made with Golang](https://img.shields.io/badge/Golang-1.21.6-blue?logo=go&logoColor=white&style=for-the-badge)](https://go.dev "Go to Golang homepage")

This Benchmark is based on the article: [Go interfaces, but at what cost? · Jul 6, 2019][article]

### Results

Here’s the result, with an Apple M1 Pro Memory 32 GB MacOS Sonoma v14.5 (23F79), using Go 1.21.6.

```bash
$ go version
go version go1.21.6 darwin/arm64

$ go test -v -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/joselitofilho/go-interface-benchmark
BenchmarkInterfaceCallSimple
BenchmarkInterfaceCallSimple/ViaInterface
BenchmarkInterfaceCallSimple/ViaInterface-10            546365151                2.076 ns/op           0 B/op          0 allocs/op
BenchmarkInterfaceCallSimple/Direct
BenchmarkInterfaceCallSimple/Direct-10                  1000000000               0.9386 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/joselitofilho/go-interface-benchmark 2.556s
```

[article]: https://syslog.ravelin.com/go-interfaces-but-at-what-cost-961e0f58a07b