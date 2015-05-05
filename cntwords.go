package read

import (
	"strings"
	"unicode"
)

// CntWords counts the number of words in a text by counting the spaces.
func CntWords(text string) int {
	cnt := 0
	wasSpace := false
	text = strings.TrimSpace(text) + " "
	for _, char := range text {
		if unicode.IsSpace(char) {
			if !wasSpace {
				cnt++
				wasSpace = true
			}
		} else {
			wasSpace = false
		}
	}
	return cnt
}
