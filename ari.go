package read

// Ari scores the Automated Readability Index test.
// See https://en.wikipedia.org/wiki/Automated_Readability_Index.
func Ari(text string) float64 {
	letterCnt, digitCnt, punctCnt := CntChars(text)
	charCnt := float64(letterCnt + digitCnt + punctCnt)
	wordCnt := float64(CntWords(text))
	sentCnt := float64(CntSents(text))
	return 4.71*(charCnt/wordCnt) + 0.5*(wordCnt/sentCnt) - 21.43

}
