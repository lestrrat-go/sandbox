package mapkey_test

import (
	"crypto/rand"
	"crypto/sha256"
	"math"
	"math/big"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var maxInt = big.NewInt(math.MaxInt64)

func randomString(min, max int) string {
	return string(randomBytes(min, max))
}

func randomBytes(min, max int) []byte {
	var n *big.Int
	if min == max {
		n = big.NewInt(int64(min))
	} else {
		x := big.NewInt(int64(max - min))
		n, _ = rand.Int(rand.Reader, x)
	}

	b := make([]byte, int(n.Int64()))
	// A randsrc.Int63() generates 63 random bits, enough for
	// letterIdxMax characters!
	cache, _ := rand.Int(rand.Reader, maxInt)
	for i, remain := int(n.Int64())-1, letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, _ = rand.Int(rand.Reader, maxInt)
			remain = letterIdxMax
		}
		if idx := int(cache.Int64() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache.SetInt64(cache.Int64() >> letterIdxBits)
		remain--
	}

	return b
}

const keyCount = 100000

func BenchmarkMapString(b *testing.B) {
	// Pre-allocate the string keys so that the cost of generating
	// string keys are not reflected in the benchmark
	keys := make([]string, keyCount)
	for i := 0; i < keyCount; i++ {
		keys[i] = randomString(8, sha256.Size)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]struct{})
		for j := 0; j < keyCount; j++ {
			m[keys[i]] = struct{}{}
		}

		for k := range m {
			_ = m[k]
		}
	}
}

func BenchmarkMapBytes(b *testing.B) {
	// Pre-allocate the [sha256.Size]byte keys so that the cost of generating
	// []byte keys are not reflected in the benchmark
	keys := make([][sha256.Size]byte, keyCount)
	for i := 0; i < keyCount; i++ {
		b := randomBytes(sha256.Size, sha256.Size)
		for j := 0; j < sha256.Size; j++ {
			keys[i][j] = b[j]
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[[sha256.Size]byte]struct{})
		for j := 0; j < keyCount; j++ {
			m[keys[j]] = struct{}{}
		}

		for k := range m {
			_ = m[k]
		}
	}
}

func BenchmarkMapInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]struct{})
		for j := 0; j < keyCount; j++ {
			m[j] = struct{}{}
		}

		for k := range m {
			_ = m[k]
		}
	}
}
