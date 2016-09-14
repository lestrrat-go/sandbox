# BENCHMARK

Check access speed between string, [fixed_size]byte, and int

```
BenchmarkMapString-4         500       3351333 ns/op           0 B/op          0 allocs/op
BenchmarkMapBytes-4           50      25427168 ns/op    10852246 B/op       6232 allocs/op
BenchmarkMapInt-4            100      15557499 ns/op     3279572 B/op       6223 allocs/op
```