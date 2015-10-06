package bufferbench

import (
	"bytes"
	"sync"
	"testing"
)

var chars []byte
func init() {
	s := "acdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	chars = make([]byte, len(s))
	for i, c := range s {
		chars[i] = byte(c)
	}
}

func BenchmarkWriteToBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := bytes.Buffer{}
		for j := 0; j < 100; j++ {
			b.Write(chars)
		}
	}
}

var pool = sync.Pool{}
func BenchmarkWriteToBufferWithPool(b *testing.B) {
	pool.New = func() interface{} {
		return &bytes.Buffer{}
	}
	for i := 0; i < b.N; i++ {
		b := pool.Get().(*bytes.Buffer)
		b.Reset()
		for j := 0; j < 100; j++ {
			b.Write(chars)
		}
		pool.Put(b)
	}
}
