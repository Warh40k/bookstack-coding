package main

import (
	"fmt"
	"github.com/Warh40k/bookstack-coding/pkg"
	"github.com/Warh40k/bookstack-coding/pkg/encoder"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: encoder <input path> <output path>")
		os.Exit(1)
	}

	inputSeq, err := pkg.GetSequence(os.Args[1])
	if err != nil {
		fmt.Errorf("error opening input file: %s", err)
		os.Exit(1)
	}
	encodedSeq := encoder.Encode(inputSeq)

	err = pkg.SaveSequence(os.Args[2], encodedSeq)
	if err != nil {
		fmt.Errorf("error creating output file: %s", err)
		os.Exit(1)
	}
}
