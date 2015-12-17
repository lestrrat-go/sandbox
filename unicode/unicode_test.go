package unicode_test

import (
	"testing"
	"unicode"
)

func TestIsVSIn(t *testing.T) {
	alphas := make([]rune, 0, 26)
	digits := make([]rune, 0, 10)
	for i := 48; i < 58; i++ {
		digits = append(digits, rune(i))
	}
	for i := 65; i < 91; i++ {
		alphas = append(alphas, rune(i))
	}
	for i := 97; i < 123; i++ {
		alphas = append(alphas, rune(i))
	}

	for _, alpha := range alphas {
		if !unicode.In(alpha, unicode.Letter, unicode.Digit) {
			t.Errorf("'%c' did not match In(Letter, Digit)", alpha)
		}
	}
	for _, digit := range digits {
		if !unicode.In(digit, unicode.Letter, unicode.Digit) {
			t.Errorf("'%c' did not match In(Letter, Digit)", digit)
		}
	}
}