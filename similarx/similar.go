package similarx

import (
	"github.com/antlabs/strsim"
)

func Compare(src, dest string, code string) float64 {
	option := strsim.UseBase64()
	var val float64
	switch code {
	//case "Levenshtein":
	//	val =  strsim.Compare(src, dest, option)
	case "Dice":
		val = strsim.Compare(src, dest, strsim.DiceCoefficient(), option)
	case "Jaro":
		val = strsim.Compare(src, dest, strsim.Jaro(), option)
	case "JaroWinkler":
		val = strsim.Compare(src, dest, strsim.JaroWinkler(), option)
	case "Hamming":
		val = strsim.Compare(src, dest, strsim.Hamming(), option)
	case "Cosine":
		val = strsim.Compare(src, dest, strsim.Cosine(), option)
	case "Simhash":
		val = strsim.Compare(src, dest, strsim.Simhash(), option)
	default:
		val = strsim.Compare(src, dest, option)
	}
	return val
}
