package read

import (
	"testing"
)

func TestCntSyls(t *testing.T) {
	exp := 10
	text := "99 bottles of b33r on the wall!!"
	act := CntSyls(text)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
