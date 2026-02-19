package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii-art/src"
)

func main() {
	fontFlag := flag.String("f", "standard", "banner ASCII font")
	flag.Parse()
	inputString := flag.Args()

	if !src.IsValidFont(*fontFlag) {
		fmt.Fprintln(os.Stderr, "Error: wrong font flag")
		fmt.Fprintln(os.Stderr, "Valid fonts: standard, shadow, thinkertoy")
		os.Exit(1)
	}
	if len(inputString) != 1 {
		fmt.Fprintln(os.Stderr, "Error: wrong arguments")
		fmt.Fprintln(os.Stderr, "Usage: go run . \"INPUT STRING\"")
		os.Exit(1)
	}
	if !src.IsValidAscii(inputString[0]) {
		fmt.Fprintln(os.Stderr, "Error: found non ascii character")
		os.Exit(1)
	}
	fontData := src.GetReader(src.FontType(*fontFlag))
	fontMap, err := src.BuildFontMap(fontData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing font map")
		os.Exit(1)
	}

	s := inputString[0]
	if s == "" {
		return
	}

	str := inputString[0]
	if str == "" {
		return
	}

	words := strings.Split(str, "\\n")
	if len(words) > 0 && words[len(words)-1] == "" {
		words = words[:len(words)-1]
	}

	for i, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			src.PrintBanner(word, fontMap)
		}
		if i < len(words)-1 {
			fmt.Println()
		}
	}

}
