# RESULTS

Benchmark run on MacBook Pro (2.9GHz Core i5, 16GB RAM), OS X 10.11.6.
Go version `go version go1.7 darwin/amd64`

```
BenchmarkPlusConcat-4           10000000           168 ns/op          35 B/op          2 allocs/op
BenchmarkFmtSprintf-4            3000000           422 ns/op          96 B/op          5 allocs/op
BenchmarkBytesBuffer-4          10000000           164 ns/op          35 B/op          2 allocs/op
BenchmarkAppendBytes-4          20000000           105 ns/op          32 B/op          1 allocs/op
BenchmarkAppendBytesNoCap-4     10000000           131 ns/op          64 B/op          2 allocs/op
```