package read

import (
	"testing"
)

func TestCntCopWords(t *testing.T) {
	exp := 2
	text := "Complexity is confounding the comprehending of everything"
	act := CntCopWords(text)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
