package read

import "unicode"

// CntChars counts the characters, numbers, and punctuation marks in a text.
func CntChars(text string) int {
	cnt := 0
	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsPunct(char) {
			cnt++
		}
	}
	return cnt
}
