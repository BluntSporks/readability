package read

import (
	"testing"
)

func TestCntPolysyls(t *testing.T) {
	exp := 3
	text := "Complexity is confounding the comprehending of everything"
	act := CntPolysyls(text)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
