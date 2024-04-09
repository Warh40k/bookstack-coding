package decoder

import (
	"bytes"
	"io"
	"slices"
)

var alph = []byte{'Z', 'Y', 'X', 'W', 'V', 'U', 'T', 'S', 'R', 'Q', 'P', 'O', 'N', 'M', 'L', 'K', 'J', 'I', 'H', 'G', 'F', 'E', 'D', 'C', 'B', 'A', ' '}

func Decode(inputSeq []byte) []byte {
	// используя алфавит, декодировать числа в исходные байты
	rngBytes := getBytesFromBinary(inputSeq)
	m := len(alph)
	result := make([]byte, len(rngBytes))
	workSeq := slices.Concat(alph, result)
	var encountered map[byte]bool

	for i := m; i < len(workSeq); i++ {
		encountered = make(map[byte]bool)
		rng := int(rngBytes[i-m])
		var j = i - 1

		for {
			if !encountered[workSeq[j]] {
				encountered[workSeq[j]] = true
				rng--
				if rng < 0 {
					break
				}
			}
			j--
		}
		workSeq[i] = workSeq[j]
		result[i-m] = workSeq[i]
	}

	return result
}

func getBytesFromBinary(inputSeq []byte) []byte {
	var seq, resBuf bytes.Buffer

	seq.Grow(len(inputSeq) * 8)
	for _, b := range inputSeq {
		writeBits(&seq, b)
	}

	var unar int
	for {
		b, err := seq.ReadByte()
		if err == io.EOF {
			break
		}

		if b == 0 {
			num := seq.Next(unar)
			resBuf.WriteByte(parseBinary(unar, num))
			unar = 0
		} else {
			unar++
		}
	}

	return resBuf.Bytes()
}

func parseBinary(n int, bin []byte) byte {
	if n == 0 {
		return 0
	}
	var result byte
	var i int
	for i = 0; i < n; i++ {
		result += 1 << i * bin[n-i-1]
	}
	result += 1 << i
	result--

	return result
}

func writeBits(seq *bytes.Buffer, b byte) {
	var bits [8]byte
	for i := 7; i >= 0; i-- {
		bits[i] = b & 1
		b >>= 1
	}
	seq.Write(bits[:])
}
