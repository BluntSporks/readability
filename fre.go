package read

// Fre scores the Flesch reading-ease.
// See https://en.wikipedia.org/wiki/Flesch–Kincaid_readability_tests#Flesch_reading_ease.
func Fre(text string) float64 {
	sylCnt := float64(CntSyls(text))
	wordCnt := float64(CntWords(text))
	sentCnt := float64(CntSents(text))	
	return 206.835 - 1.015*(wordCnt/sentCnt) - 84.6*(sylCnt/wordCnt)
}
