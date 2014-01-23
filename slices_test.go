package main

import(
  "testing"
)

/*

BenchmarkSliceAppend(), BenchmarkSliceSet(), and BenchmarkArraySet() compare
the various ways in which we create an array / slice from somewhere else.

In my tests:

  $ go test -benchtime 5s -bench .
  BenchmarkSliceAppend  2000000000           1.46 ns/op
  BenchmarkSliceSet     1000000000           0.41 ns/op
  BenchmarkArraySet     2000000000           0.15 ns/op

So it's always better to allocate the require length and then set

*/

const SLICE_SIZE = 10000


func BenchmarkSliceAppend(t *testing.B) {
  var array [SLICE_SIZE]int
  for i := 1; i <= SLICE_SIZE; i++ {
    array[i - 1] = i
  }

  for i := 1; i <= SLICE_SIZE; i++ {
    var slice []int
    for _, x := range array {
      slice = append(slice, x)
    }

    // Sanity check
    for idx, x := range slice {
      if idx + 1 != x {
        t.Fatalf("contents do not match (index %d != value %d)", idx + 1, x)
      }
    }
    if len(slice) != SLICE_SIZE {
      t.Fatalf("slice size does not match (got %d, want %d)", len(slice), SLICE_SIZE)
    }
  }
}

func BenchmarkSliceSet(t *testing.B) {
  var array [SLICE_SIZE]int
  for i := 1; i <= SLICE_SIZE; i++ {
    array[i - 1] = i
  }

  for i := 1; i <= SLICE_SIZE; i++ {
    slice := make([]int, SLICE_SIZE)
    for j, x := range array {
      slice [j] = x
    }

    // Sanity check
    for idx, x := range slice {
      if idx + 1 != x {
        t.Fatalf("contents do not match (index %d != value %d)", idx + 1, x)
      }
    }
    if len(slice) != SLICE_SIZE {
      t.Fatalf("slice size does not match (got %d, want %d)", len(slice), SLICE_SIZE)
    }
  }
}

func BenchmarkArraySet(t *testing.B) {
  var array [SLICE_SIZE]int
  for i := 1; i <= SLICE_SIZE; i++ {
    array[i - 1] = i
  }

  for i := 1; i <= SLICE_SIZE; i++ {
    var arrayCopy [SLICE_SIZE]int
    for j, x := range array {
      arrayCopy[j] = x
    }

    // Sanity check
    for idx, x := range arrayCopy {
      if idx + 1 != x {
        t.Fatalf("contents do not match (index %d != value %d)", idx + 1, x)
      }
    }
    if len(arrayCopy) != SLICE_SIZE {
      t.Fatalf("arrayCopy size does not match (got %d, want %d)", len(arrayCopy), SLICE_SIZE)
    }
  }
}
