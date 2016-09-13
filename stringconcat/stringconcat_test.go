package stringconcat

import (
	"bytes"
	"fmt"
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
}

func BenchmarkFmtSprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%s%s%d",
			"abcdefghijklmnopqrstuvwxyz",
			[]byte{'a', 'b', 'c'},
			100,
		)
		_ = s
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteString("abcdefghijklmnopqrstuvwxyz")
		buf.Write([]byte{'a', 'b', 'c'})
		buf.WriteString(strconv.Itoa(100))
	}
}
