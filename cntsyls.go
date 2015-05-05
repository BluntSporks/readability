package read

// CntSyls estimates the syllable counts in a text from the number of characters in words.
func CntSyls(text string) int {
	letterCnt, digitCnt, _ := CntChars(text)
	return letterCnt/3 + digitCnt
}
