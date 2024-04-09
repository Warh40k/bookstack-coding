package encoder

import (
	"bytes"
	"fmt"
	"github.com/Warh40k/bookstack-coding/pkg"
	"math"
	"slices"
	"strings"
)

//var alph = []byte{'Z', 'Y', 'X', 'W', 'V', 'U', 'T', 'S', 'R', 'Q', 'P', 'O', 'N', 'M', 'L', 'K', 'J', 'I', 'H', 'G', 'F', 'E', 'D', 'C', 'B', 'A', ' '}

func Encode(inputSeq []byte) []byte {
	alph := pkg.GetAlphabet()
	//slices.Reverse(alph)
	m := len(alph)
	workingSeq := slices.Concat(alph, inputSeq)
	bitstr := bytes.Buffer{}
	for i := m; i < len(workingSeq); i++ {
		var rng int
		var encountered = make(map[byte]bool)
		for j := i - 1; j >= 0; j-- {
			if workingSeq[j] == workingSeq[i] {
				break
			}
			if !encountered[workingSeq[j]] {
				encountered[workingSeq[j]] = true
				rng++
			}
		}
		encodedSym := GetUnar(rng)
		bitstr.WriteString(encodedSym)
	}

	result := convertToBytes(&bitstr)

	return result
}

func GetUnar(rng int) string {
	rng += 2 // прибавляем единицу (чтобы не брать log2 от 0 и 1)
	bcount := int(math.Log2(float64(rng)))
	bin := fmt.Sprintf("%b", rng)[1:] // отсекаем старший бит (он всегда равен 1)
	return fmt.Sprintf(strings.Repeat("1", bcount-1)+"0"+"%s", bin)
}

func convertToBytes(bitstr *bytes.Buffer) []byte {
	buflen := bitstr.Len()
	var result = make([]byte, int(math.Ceil(float64(buflen)/8)))
	var bufstring = bitstr.String()

	for i := 0; i < buflen; i += 8 {
		var b byte
		var j int

		for j = 0; j < 8; j++ {
			b = b << 1
			if i+j >= buflen {
				b |= 1
				continue
			}
			if bufstring[i+j] == '1' {
				b |= 1
			}
		}
		result[i/8] = b
	}
	return result
}
