package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

const ALPHABET_SIZE = 256

// GetSequence получить входные данные из файла
func GetSequence(inputPath string) ([]rune, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var inputSeq = make([]rune, 0)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		inputSeq = append(inputSeq, r)
	}
	return inputSeq, nil
}

// SaveSequence результат в файл
func SaveSequence(outputPath string, outputSeq []byte) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.Write(outputSeq)
	writer.Flush()

	return nil
}

func GetAlphabet() []rune {
	var alph = make([]rune, ALPHABET_SIZE)
	var i rune
	for i = ALPHABET_SIZE - 1; i >= 0; i-- {
		alph[i] = i
	}
	return alph
}

func GetUnar(rng int) string {
	bcount := int(math.Log2(float64(rng+1))) + 1
	binrng := []rune(fmt.Sprintf("%b", rng))[1:]
	return fmt.Sprintf(strings.Repeat("1", bcount-1)+"0"+"%s", string(binrng))
}

func ConvertToBytes(bitstr bytes.Buffer) []byte {
	buflen := bitstr.Len()
	var result = make([]byte, int(math.Ceil(float64(buflen/8))))
	var bufstring = bitstr.String()
	for i := 0; i < buflen-8; i += 8 {
		var b byte

		for j := 0; j < 8; j++ {
			b = b << 1
			if i+j >= buflen {
				continue
			}
			if bufstring[i+j] == '1' {
				b = b | 1
			}
		}
		result[i/8] = b
	}
	return result
}
