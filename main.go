package main

import (
	"ascii-art/fonts"
	"ascii-art/src"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fontFlag := flag.String("f", "standard", "banner ASCII font")
	flag.Parse()
	inputString := flag.Args()

	if !fonts.IsValidFont(*fontFlag) {
		fmt.Fprintln(os.Stderr, "Error: wrong font flag")
		fmt.Fprintln(os.Stderr, "Valid fonts: standard, shadow, thinkertoy")
		os.Exit(1)
	}
	if len(inputString) != 1 {
		fmt.Fprintln(os.Stderr, "Error: wrong arguments")
		fmt.Fprintln(os.Stderr, "Usage: go run . \"INPUT STRING\"")
		os.Exit(1)
	}
	fontData := fonts.GetReader(fonts.FontType(*fontFlag))
	fontMap, err := src.BuildFontMap(fontData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing font map")
		os.Exit(1)
	}

	// split the user input by \n and print the words
	words := strings.SplitSeq(inputString[0], "\\n")
	for word := range words {
		if word == "" {
			fmt.Println()
			continue
		}
		src.PrintBanner(word, fontMap)
	}
}
