package stringconcat

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"testing"
)

var s1 = "abcdefghijklmnopqrstuvwxyz"
var b1 = []byte{'a', 'b', 'c'}
var i1 = 100

func BenchmarkPlusConcat(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ""
		s = s + s1 + string(b1) + strconv.Itoa(i1)
		io.WriteString(ioutil.Discard, s)
	}
}

func BenchmarkFmtSprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%s%s%d", s1, b1, i1)
		io.WriteString(ioutil.Discard, s)
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteString(s1)
		buf.Write(b1)
		buf.WriteString(strconv.Itoa(i1))
		io.WriteString(ioutil.Discard, buf.String())
	}
}

func BenchmarkAppendBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 0, 64)
		buf = append(buf, s1...)
		buf = append(buf, b1...)
		buf = strconv.AppendInt(buf, int64(i1), 10)
		io.WriteString(ioutil.Discard, string(buf))
	}
}

// same as BenchmarkAppendBytes, but without pre-allocating
// the buffer
func BenchmarkAppendBytesNoCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf []byte
		buf = append(buf, s1...)
		buf = append(buf, b1...)
		buf = strconv.AppendInt(buf, int64(i1), 10)
		io.WriteString(ioutil.Discard, string(buf))
	}
}
