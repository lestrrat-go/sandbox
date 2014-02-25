package generator

// This shows why the original benchmark is useless.
// http://jxck.hatenablog.com/entry/go-generator-benchmark
//
// When run with `go test -cpu 4 ...`, the closure version
// barfs, as it does not protect from race conditions. The
// channel version is safe, as channel communication does all
// the locking underneath.
//
// It's because of these locks that channels seem "slow".
// The original benchmark only uses 1 method in 1 goroutine,
// so it's comparison is useless.
//
// Sample output on my OS X 10.9.1, go 1.2
// yamaneko:generator daisuke$ go test -cpu 4 -benchtime 2s -bench .
// PASS
// BenchmarkChannel-4  10000000         600 ns/op
// BenchmarkClosure-4  1000000000           3.56 ns/op
// BenchmarkChannelGoroutine-4  5000000         861 ns/op
// BenchmarkClosureGoroutine-4 panic: Key 34 already exists in map!
//
// goroutine 69 [running]:
// [stacktrace omitted]

import (
  "fmt"
  "sync"
  "testing"
)

func TestChannel(t *testing.T) {
  ch := Channel()

  for i := 1; i < 100; i++ {
    actual := <-ch
    expected := i
    if actual != expected {
      t.Errorf("\ngot  %v\nwant %v", actual, expected)
    }
  }
}

func TestClosure(t *testing.T) {
  cl := Closure()

  for i := 1; i < 100; i++ {
    actual := cl()
    expected := i
    if actual != expected {
      t.Errorf("\ngot  %v\nwant %v", actual, expected)
    }
  }
}

func BenchmarkChannel(b *testing.B) {
  ch := Channel()
  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    <-ch
  }
}

func BenchmarkClosure(b *testing.B) {
  cl := Closure()
  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    cl()
  }
}

func BenchmarkChannelGoroutine(b *testing.B) {
  ch := Channel()
  b.ResetTimer()

  n_goro := 10
  n_per_goro := int(b.N / n_goro)

  values := make([]map[int]bool, n_goro)
  for i := 0; i < n_goro; i++ {
    values[i] = make(map[int]bool)
  }

  wg := &sync.WaitGroup {}
  for i := 0; i < n_goro; i++ {
    wg.Add(1)
    local_i := i
    go func() {
      defer wg.Done()
      for j := 0; j < n_per_goro; j++ {
        k := <-ch
        if _, exists := values[local_i][k]; exists {
          panic(fmt.Sprintf("Key %d already exists in map!", k))
        }
        values[local_i][k] = true
      }
    }()
  }
  wg.Wait()
  merged := make(map[int]bool)
  for i := 0; i < n_goro; i++ {
    for k, _ := range values[i] {
      if _, exists := merged[k]; exists {
        panic(fmt.Sprintf("Key %d already exists in map!", k))
      }
      merged[k] = true
    }
  }
}

func BenchmarkClosureGoroutine(b *testing.B) {
  cl := Closure()
  b.ResetTimer()

  n_goro := 10
  n_per_goro := int(b.N / n_goro)
  values := make([]map[int]bool, n_goro)
  for i := 0; i < n_goro; i++ {
    values[i] = make(map[int]bool)
  }
  wg := &sync.WaitGroup {}
  for i := 0; i < n_goro; i++ {
    wg.Add(1)
    local_i := i
    go func() {
      defer wg.Done()
      for j := 0; j < n_per_goro; j++ {
        k := cl()
        if _, exists := values[local_i][k]; exists {
          panic(fmt.Sprintf("Key %d already exists in map!", k))
        }
        values[local_i][k] = true
      }
    }()
  }

  wg.Wait()
  merged := make(map[int]bool)
  for i := 0; i < n_goro; i++ {
    for k, _ := range values[i] {
      if _, exists := merged[k]; exists {
        panic(fmt.Sprintf("Key %d already exists in map!", k))
      }
      merged[k] = true
    }
  }
}
