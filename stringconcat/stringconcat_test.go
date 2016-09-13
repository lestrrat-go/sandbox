package stringconcat

import (
	"bytes"
	"strconv"
	"testing"
)

func BenchmarkPlusConcat(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ""
		s = s +
			"abcdefghijklmnopqrstuvwxyz" +
			string([]byte{'a', 'b', 'c'}) +
			strconv.Itoa(100)
	}
	b.Logf(s)
}

func BenchmarkBytesBuffer(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteString("abcdefghijklmnopqrstuvwxyz")
		buf.Write([]byte{'a', 'b', 'c'})
		buf.WriteString(strconv.Itoa(100))
	}
	b.Logf(buf.String())
}
