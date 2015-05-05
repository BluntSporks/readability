package read

import (
	"testing"
)

func TestCntWords(t *testing.T) {
	exp := 6
	text := "  Space ...   the final	: frontier "
	act := CntWords(text)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
