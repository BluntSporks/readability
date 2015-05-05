package read

import "math"

// Smog scores the SMOG score.
// See https://en.wikipedia.org/wiki/SMOG.
func Smog(text string) float64 {
	polysylCnt := float64(CntPolysyls(text))
	sentCnt := float64(CntSents(text))
	return 1.0430*math.Sqrt(polysylCnt*30.0/sentCnt) + 3.1291

}
