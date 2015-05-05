package read

import (
	"testing"
)

func TestCntSents(t *testing.T) {
	exp := 4
	text := " This.1 'Is.' only a ? Test!! of.sentence.counts "
	act := CntSents(text)
	if exp != act {
		t.Error("Expected", exp, "got", act)
	}
}
