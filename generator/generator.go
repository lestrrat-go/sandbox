package generator

// http://jxck.hatenablog.com/entry/go-generator-benchmark

func Channel() <-chan int {
  i := 1
  ch := make(chan int)
  go func() {
    for {
      ch <- i
      i = i + 1
    }
  }()
  return ch
}

func Closure() func() int {
  i := 0
  return func() int {
    i = i + 1
    return i
  }
}