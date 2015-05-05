package read

import (
	"unicode"
)

// CntCopWords counts the number of complex words in a text.
// This is an attempt to define the notion of complex like Gunning Fog but with simple computation.
// See https://en.wikipedia.org/wiki/Gunning_fog_index.
func CntCopWords(text string) int {
	cnt := 0
	word := ""
	runeLen := 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			word += string(char)
			runeLen++
		} else if word != "" {
			// Adjust rune length for common endings.
			if runeLen >= 9 {
				end2 := word[len(word)-2:]
				if end2 == "es" || end2 == "ed" {
					runeLen -= 2
				} else {
					end3 := word[len(word)-3:]
					if end3 == "ing" {
						runeLen -= 3
					}
				}
			}

			// If rune length is still >= 9, the word is complex.
			if runeLen >= 9 {
				cnt++
			}
			runeLen = 0
		}
	}
	return cnt
}
