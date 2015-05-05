package read

// Fk scores the Flesch-Kincaid Grade Level.
// See https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests.
func Fk(text string) float64 {
	sylCnt := float64(CntSyls(text))
	wordCnt := float64(CntWords(text))
	sentCnt := float64(CntSents(text))
	return 0.39*(wordCnt/sentCnt) + 11.8*(sylCnt/wordCnt) - 15.59

}
