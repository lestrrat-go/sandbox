package unicode_test

import (
	"testing"
	"unicode"
)

var (
	alphas []rune
	digits []rune
)

func init() {
	alphas = make([]rune, 0, 31)
	digits = make([]rune, 0, 20)
	for i := 48; i < 58; i++ {
		digits = append(digits, rune(i))
	}
	digits = append(digits, '１', '２', '３', '４', '５', '６', '７', '８', '９', '０')

	for i := 65; i < 91; i++ {
		alphas = append(alphas, rune(i))
	}
	for i := 97; i < 123; i++ {
		alphas = append(alphas, rune(i))
	}
	alphas = append(alphas, 'あ', 'い', 'う', 'え', 'お')
}

func BenchmarkIs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, alpha := range alphas {
			if !unicode.IsLetter(alpha) && !unicode.IsDigit(alpha) {
				b.Errorf("'%c' did not match In(Letter, Digit)", alpha)
			}
		}
		for _, digit := range digits {
			if !unicode.IsLetter(digit) && !unicode.IsDigit(digit) {
				b.Errorf("'%c' did not match In(Letter, Digit)", digit)
			}
		}
	}
}

func IsNmchar(r rune) bool {
	if uint32(r) <= unicode.MaxLatin1 {
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	}
	return unicode.In(r, unicode.Letter, unicode.Digit)
}

func BenchmarkIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, alpha := range alphas {
			if !IsNmchar(alpha) {
				b.Errorf("'%c' did not match In(Letter, Digit)", alpha)
			}
		}
		for _, digit := range digits {
			if !IsNmchar(digit) {
				b.Errorf("'%c' did not match In(Letter, Digit)", digit)
			}
		}
	}
}
