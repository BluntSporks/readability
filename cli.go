package read

// Cli scores the Coleman-Liau Index.
// See https://en.wikipedia.org/wiki/Coleman%E2%80%93Liau_index.
func Cli(text string) float64 {
	letterCnt, digitCnt, _ := CntChars(text)
	charCnt := float64(letterCnt + digitCnt)
	wordCnt := float64(CntWords(text))
	sentCnt := float64(CntSents(text))

	// Calculate letters per 100 words.
	lp100w := charCnt / wordCnt * 100.0

	// Calculate sentences per 100 words.
	sp100w := sentCnt / wordCnt * 100.0

	// Return CLI.
	return 0.0588*lp100w - 0.296*sp100w - 15.8

}
