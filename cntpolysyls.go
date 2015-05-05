package read

import (
	"unicode"
)

// CntPolysyls counts the number of polysyllable words in a text.
func CntPolysyls(text string) int {
	cnt := 0
	word := ""
	runeLen := 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			word += string(char)
			runeLen++
		} else if word != "" {
			// If rune length is >= 9, the word is likely a polysyllable.
			if runeLen >= 9 {
				cnt++
			}
			runeLen = 0
		}
	}
	return cnt
}
