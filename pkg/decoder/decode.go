package decoder

import (
	"bytes"
	"fmt"
	"io"
)

func Decode(inputSeq []byte) []byte {
	// вытащить битовое каждого байта
	// пробежаться по полученной строке
	// извлекать унарный код и декодировать следующие n бит числа
	// используя алфавит, декодировать числа в исходные байты
	var seq = bytes.Buffer{}

	seq.Grow(len(inputSeq) * 8)
	for _, b := range inputSeq {
		writeBits(&seq, b)
	}
	//fmt.Println(seq.Bytes())
	var resBytes bytes.Buffer
	var unar int
	for {
		b, err := seq.ReadByte()
		if err == io.EOF {
			break
		}

		if b == 0 {
			//n-- // учет манипуляций энкодера
			num := seq.Next(unar)
			resBytes.WriteByte(parseBinary(unar, num))
			unar = 0
		} else {
			unar++
		}
	}

	fmt.Println(resBytes.Bytes())
	return nil
}

/*func parseUnar(n int, unar []byte) int {
	var result byte
	for i := n; i >= 0; i-- {
		result += 1 << i * unar[i]
	}
	return int(result)
}*/

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
