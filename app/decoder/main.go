package main

import (
	"fmt"
	"github.com/Warh40k/bookstack-coding/pkg"
	"github.com/Warh40k/bookstack-coding/pkg/decoder"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: decoder <input path> <output path>")
		os.Exit(1)
	}

	inputSeq, err := pkg.GetSequence(os.Args[1])
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
}
