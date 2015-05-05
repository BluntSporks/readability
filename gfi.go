package read

// Gfi scores the Gunning fog index.
// See https://en.wikipedia.org/wiki/Gunning_fog_index.
func Gfi(text string) float64 {
	wordCnt := float64(CntWords(text))
	sentCnt := float64(CntSents(text))
	copWordCnt := float64(CntCopWords(text))
	return 0.4 * (wordCnt/sentCnt + 100.0*copWordCnt/wordCnt)

}
