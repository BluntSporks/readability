package read

// Fre German
// Fre scores the Flesch reading-ease. (German)
// See https://de.wikipedia.org/wiki/Lesbarkeitsindex
func FreGerman(text string) float64 {
	sylCnt := float64(read.CntSyls(text))
	wordCnt := float64(read.CntWords(text))
	sentCnt := float64(read.CntSents(text))
	score := 180 - (wordCnt / sentCnt) - 58.5*(sylCnt/wordCnt)
	return score
}
