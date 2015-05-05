package read

import (
	"strings"
	"unicode"
)

// CntSents counts the number of sentences in a text by counting ending marks.
func CntSents(text string) int {
	cnt := 0
	wasEnd := false
	text = strings.TrimSpace(text) + " "
	for _, char := range text {
		if char == '.' || char == '?' || char == '!' {
			wasEnd = true
		} else if unicode.IsLetter(char) {
			wasEnd = false
		} else if wasEnd && unicode.IsSpace(char) {
			cnt++
			wasEnd = false
		}
	}
	return cnt
}
