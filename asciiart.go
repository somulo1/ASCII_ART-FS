package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("pass 3 arguments i.e usage: go run . hello thinkertoy.txt")
		return
	}
	inputFile := os.Args[2]
	// read the banner file
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var fileLines []string

	if inputFile == "thinkertoy.txt" {
		fileLines = strings.Split(string(fileContent), "\r\n")
	} else {
		fileLines = strings.Split(string(fileContent), "\n")
	}

	input := os.Args[1]
	if input == "" {
		return
	}

	// Split input by newline escape sequence
	words := strings.Split(input, "\\n")
	for _, word := range words {
		for _, char := range word {
			// To check for non-ascii
			if !(char >= ' ' && char <= '~') {
				fmt.Println("Non-ascii detected!! Check and try again")
				return
			}
			for i := 0; i < 8; i++ {
				for _, char := range word {
					// Print the corresponding line from the banner content
					fmt.Print(fileLines[i+(int(char-' ')*9)+1])
				}
				fmt.Println()
			}

		}
	}
}
