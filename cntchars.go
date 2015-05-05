package read

import "unicode"

// CntChars counts the letters, digits, and punctuation marks in a text.
func CntChars(text string) (int, int, int) {
	letterCnt := 0
	digitCnt := 0
	punctCnt := 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			letterCnt++
		} else if unicode.IsDigit(char) {
			digitCnt++
		} else if unicode.IsPunct(char) {
			punctCnt++
		}
	}
	return letterCnt, digitCnt, punctCnt
}
