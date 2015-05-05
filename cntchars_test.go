package read

import (
	"testing"
)

func TestCntChars(t *testing.T) {
	expLetterCnt := 20
	expDigitCnt := 4
	expPunctCnt := 3
	text := "99 bottles, of b33r on the wall!!"
	actLetterCnt, actDigitCnt, actPunctCnt := CntChars(text)
	if expLetterCnt != actLetterCnt {
		t.Error("Expected", expLetterCnt, "got", actLetterCnt)
	}
	if expDigitCnt != actDigitCnt {
		t.Error("Expected", expDigitCnt, "got", actDigitCnt)
	}
	if expPunctCnt != actPunctCnt {
		t.Error("Expected", expPunctCnt, "got", actPunctCnt)
	}
}
