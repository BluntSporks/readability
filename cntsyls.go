package read

import "unicode"

// CntSyls estimates the syllable counts in a text from the number of characters in words.
func CntSyls(text string) int {
	letterCnt := 0
	digitCnt := 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			letterCnt++
		} else if unicode.IsDigit(char) {
			digitCnt++
		}
	}
	return letterCnt/3 + digitCnt
}
