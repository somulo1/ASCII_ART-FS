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

	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	err = FileContentcheck(fileContent)
	if err != nil {
		fmt.Println("Error:", err)
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
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				// Print the corresponding line from the file content
				fmt.Print(fileLines[i+(int(letter-' ')*9)+1])
			}
			fmt.Println()
		}
	}
}

// FileContentcheck checks the validity of the file content
func FileContentcheck(content []byte) error {
	var fileLines []string
	for _, contents := range content {
		if contents < ' ' || contents > '~' {
			fmt.Println("non-ASCII character detected in file content")
			os.Exit(0)
		}
	}
	if len(fileLines) != 856 {
		fmt.Println("file content has been tampered with")
		os.Exit(0)
	}
	return nil
}
