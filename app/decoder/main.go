package main

import (
	"fmt"
	"github.com/Warh40k/bookstack-coding/pkg"
	"github.com/Warh40k/bookstack-coding/pkg/decoder"
	"io/fs"
	"os"
	"path/filepath"
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
	var files []string
	var dirs []string
	if info.IsDir() {
		isDir = true
		err = filepath.WalkDir(os.Args[1], func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				dirs = append(dirs, path)
			} else {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("error getting dir files: %s\n", err)
			os.Exit(1)
		}
	} else {
		files = append(files, os.Args[1])
	}

	info, err = os.Stat(os.Args[2])
	if isDir && os.IsNotExist(err) {
		err = os.MkdirAll(os.Args[2], 0777)
		if err != nil {
			fmt.Printf("error creating output directory: %s\n", err)
			os.Exit(1)
		}
	}

	for _, file := range files {
		go func() {
			inputSeq, err := pkg.GetSequence(file)
			if err != nil {
				fmt.Printf("error opening input file: %s\n", err)
				os.Exit(1)
			}
			encodedSeq := decoder.Decode(inputSeq)

			err = pkg.SaveSequence(os.Args[2], encodedSeq)
			if err != nil {
				fmt.Printf("error creating output file: %s\n", err)
				os.Exit(1)
			}
		}()

	}

}
