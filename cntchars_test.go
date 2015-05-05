package read

import (
	"testing"
)

func TestCntChars(t *testing.T) {
	exp := 26
	text := "99 bottles of b33r on the wall!!"
	act := CntChars(text)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
