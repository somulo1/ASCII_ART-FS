package functions

import (
	"fmt"
	"strings"
)

func AsciiArt(stringInput string, fileLine []string) string {
	result := ""

	// replacing every instance of new line with the newline character (\n)
	stringInput = strings.Replace(stringInput, "\n", "\\n", -1)

	if !ValidSentence(stringInput) {
		return ""
	}

	// slicing the input based on the presence of the string "\n"
	words := strings.Split(stringInput, "\\n")

	empty := EmptyArray(words)
	if empty != "false" {
		fmt.Print(empty)
		return ""
	}

	for _, word := range words {
		if word == "" {
			result += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(word); j++ {
					start := (int(word[j]-' ') * 9) + 1 // calculating the begining of a character based on data from standard.txt

					result += fileLine[start+i]
				}
				result += "\n"
			}
		}
	}
	return result
}

func ValidSentence(word string) bool {
	for _, letter := range word {
		if !(letter >= ' ' && letter <= '~') {
			fmt.Println("Error, character", string(letter), "is an invalid character!!!!")
			return false
		}
	}
	return true
}

func EmptyArray(words []string) string {
	result := ""

	for i, word := range words {
		if word != "" {
			result = "false"
			return result
		}
		if i != 0 {
			result += "\n"
		}
	}
	return result
}
