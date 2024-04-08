package encoder

import (
	"bytes"
	"github.com/Warh40k/bookstack-coding/pkg"
	"slices"
)

//var alph = []rune{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a',
//
//	'Z', 'Y', 'X', 'W', 'V', 'U', 'T', 'S', 'R', 'Q', 'P', 'O', 'N', 'M', 'L', 'K', 'J', 'I', 'H', 'G', 'F', 'E', 'D', 'C', 'B', 'A', ' '}

func Encode(inputSeq []rune) []byte {
	alph := pkg.GetAlphabet()
	slices.Reverse(alph)
	m := len(alph)
	workingSeq := slices.Concat(alph, inputSeq)
	bitstr := bytes.Buffer{}
	for i := m; i < len(workingSeq); i++ {
		var rng int
		var encountered = make(map[rune]bool)
		for j := i - 1; j >= 0; j-- {
			if workingSeq[j] == workingSeq[i] {
				break
			}
			if !encountered[workingSeq[j]] {
				encountered[workingSeq[j]] = true
				rng++
			}
		}
		encodedSym := pkg.GetUnar(rng)
		bitstr.WriteString(encodedSym)
	}

	result := pkg.ConvertToBytes(bitstr)

	return result
}
