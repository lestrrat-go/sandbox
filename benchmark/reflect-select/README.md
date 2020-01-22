Reflect Select vs Raw Select

```
finch% go test -v -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/lestrrat-go/sandbox/benchmark/reflect-select
BenchmarkReflectSelect-4   	 1000000	      1189 ns/op	     266 B/op	       6 allocs/op
BenchmarkRawSelect-4       	 1781895	       670 ns/op	       4 B/op	       0 allocs/op
PASS
ok  	github.com/lestrrat-go/sandbox/benchmark/reflect-select	3.086s
```
