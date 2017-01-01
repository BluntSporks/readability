package read

// Fre scores the Flesch reading-ease.
// See https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FFlesch%E2%80%93Kincaid_readability_tests%23Flesch_reading_ease.
func Fre(text string) float64 {
	sylCnt := float64(CntSyls(text))
	wordCnt := float64(CntWords(text))
	sentCnt := float64(CntSents(text))	
	return 206.835 - 1.015*(wordCnt/sentCnt) - 84.6*(sylCnt/wordCnt)
}
