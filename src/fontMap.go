package src

import (
	"bufio"
	"io"
)

func BuildFontMap(fontByte io.Reader) (map[rune][]string, error) {
	fontMap := make(map[rune][]string)
	b := bufio.NewScanner(fontByte)

	b.Scan()

	// Starting with space (32)
	currentChar := rune(32)

	for {
		var lines []string
		for i := 0; i < 8; i++ {
			if !b.Scan() {
				return fontMap, nil
			}
			lines = append(lines, b.Text())
		}
		fontMap[currentChar] = lines
		currentChar++
		// end of ascii character
		if !b.Scan() {
			break
		}
	}
	return fontMap, b.Err()
}
