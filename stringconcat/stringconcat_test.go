package stringconcat

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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
		io.WriteString(ioutil.Discard, s)
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
		io.WriteString(ioutil.Discard, s)
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteString("abcdefghijklmnopqrstuvwxyz")
		buf.Write([]byte{'a', 'b', 'c'})
		buf.WriteString(strconv.Itoa(100))
		io.WriteString(ioutil.Discard, buf.String())
	}
}

func BenchmarkAppendBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 0, 64)
		buf = append(buf, "abcdefghijklmnopqrstuvwxyz"...)
		buf = append(buf, []byte{'a', 'b', 'c'}...)
		buf = strconv.AppendInt(buf, 100, 10)
		io.WriteString(ioutil.Discard, string(buf))
	}
}

// same as BenchmarkAppendBytes, but without pre-allocating
// the buffer
func BenchmarkAppendBytesNoCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf []byte
		buf = append(buf, "abcdefghijklmnopqrstuvwxyz"...)
		buf = append(buf, []byte{'a', 'b', 'c'}...)
		buf = strconv.AppendInt(buf, 100, 10)
		io.WriteString(ioutil.Discard, string(buf))
	}
}
