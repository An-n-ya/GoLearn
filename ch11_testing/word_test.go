package ch11_testing

import (
	"goLearn/ch11_testing/word"
	"testing"
)

func TestPalindrome(t *testing.T) {
	if !word.IsPalindrome("abcccba") {
		t.Error(`IsPalindrome("adcccba") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if word.IsPalindrome("word") {
		t.Error(`IsPalindrome("word") = true`)
	}
}
