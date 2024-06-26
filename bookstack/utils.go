package bookstack

import (
	"bufio"
	"io"
	"os"
)

const ALPHABET_SIZE = 256

// GetSequence получить входные данные из файла
func GetSequence(inputPath string) ([]byte, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var inputSeq = make([]byte, 0)
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		inputSeq = append(inputSeq, b)
	}
	return inputSeq, nil
}

// SaveSequence результат в файл
func SaveSequence(outputPath string, outputSeq []byte) (os.FileInfo, error) {
	file, err := os.Create(outputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	file.Write(outputSeq)
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return info, nil
}

func GetAlphabet() []byte {
	var alph = make([]byte, ALPHABET_SIZE)
	var i int
	for i = ALPHABET_SIZE - 1; i >= 0; i-- {
		alph[i] = byte(i)
	}
	return alph
}
