package src

import "fmt"

func PrintBanner(inputString string, fontMap map[rune][]string) {
	// Loop through each of the 8 lines of height
	// for each line, iterate through every character in the user's string
	// check if the character is in our font ascii file
	for line := 0; line < 8; line++ {
		for _, rn := range inputString {
			if ascii, exists := fontMap[rn]; exists {
				fmt.Print(ascii[line])
			}
		}
		fmt.Println()
	}
}
