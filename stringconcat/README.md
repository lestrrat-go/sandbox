# RESULTS

Benchmark run on MacBook Pro (2.9GHz Core i5, 16GB RAM), OS X 10.11.6.
Go version `go version go1.7 darwin/amd64`

```
BenchmarkPlusConcat-4           10000000           120 ns/op          35 B/op          2 allocs/op
BenchmarkFmtSprintf-4            3000000           406 ns/op          96 B/op          5 allocs/op
BenchmarkBytesBuffer-4          20000000            87.2 ns/op         3 B/op          1 allocs/op
BenchmarkApeendBytes-4          50000000            30.0 ns/op         0 B/op          0 allocs/op
BenchmarkApeendBytesNoCap-4     20000000            68.3 ns/op        32 B/op          1 allocs/op
```