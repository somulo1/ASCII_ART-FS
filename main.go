package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-fs/functions"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) < 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}

	stringInput := os.Args[1] // Reading the string argument entered

	if stringInput == "" {
		return
	}

	BannerFile := "standard.txt"

	if len(os.Args) == 3 {
		banner := strings.Replace(os.Args[2], ".txt", "", 1)
		BannerFile = banner + ".txt"
	}

	// read banner file specified
	file, err := os.ReadFile(BannerFile)
	if err != nil {
		fmt.Println("Error openning", BannerFile, err)
		return
	}
	file = []byte(strings.Replace(string(file), "\r\n", "\n", -1))
	var fileLine []string

	fileLine = strings.Split(string(file), "\n")

	link := ""
	switch BannerFile {
	case "standard.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt"
	case "shadow.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt"
	case "thinkertoy.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt"
	}

	if len(fileLine) != 856 {
		fmt.Println("The file", BannerFile, "is not correctly formated, please use the correct version", link, "!!!")
		return
	}

	fmt.Print(functions.AsciiArt(stringInput, fileLine))
}
