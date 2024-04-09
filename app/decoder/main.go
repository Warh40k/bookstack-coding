package main

import (
	"fmt"
	"github.com/Warh40k/bookstack-coding/pkg"
	"github.com/Warh40k/bookstack-coding/pkg/decoder"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var isDir bool

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: decoder <input path> <output path>")
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Printf("error opening input file: %s\n", err)
		os.Exit(1)
	}
	var inFiles []string
	var dirs []string
	if info.IsDir() {
		isDir = true
		err = filepath.WalkDir(os.Args[1], func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				dirs = append(dirs, path)
			} else {
				inFiles = append(inFiles, path)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("error getting dir files: %s\n", err)
			os.Exit(1)
		}
	} else {
		inFiles = append(inFiles, os.Args[1])
	}

	info, err = os.Stat(os.Args[2])
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(os.Args[2]), 0777)
		if err != nil {
			fmt.Printf("error creating output directory: %s\n", err)
			os.Exit(1)
		}
	}
	for _, dir := range dirs {
		err = os.MkdirAll(filepath.Join(os.Args[2], strings.Split(dir, os.Args[1])[1]), 0777)
		if err != nil {
			fmt.Printf("error creating output directory: %s\n", err)
			os.Exit(1)
		}
	}

	for i := 0; i < len(inFiles); i++ {
		inputSeq, err := pkg.GetSequence(inFiles[i])
		if err != nil {
			fmt.Printf("error opening input file: %s\n", err)
			os.Exit(1)
		}
		encodedSeq := decoder.Decode(inputSeq)

		var outPath = os.Args[2]
		if isDir {
			outPath = strings.Replace(inFiles[i], os.Args[1], os.Args[2], 1)
		}

		err = pkg.SaveSequence(outPath, encodedSeq)
		if err != nil {
			fmt.Printf("error creating output file: %s\n", err)
			os.Exit(1)
		}
	}

}
