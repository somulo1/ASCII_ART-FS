package main

import (
	"ascii-art-fs/functions"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAsciiArtStandard(t *testing.T) {

	inputFile := "standard.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var fileLine []string

	fileLine = strings.Split(string(file), "\n")

	//create a map containing the input text as key and the path to its expected output as value
	test_cases := map[string]string{
		"hello":         "test_cases/expectedOutput1.txt",
		"you & me":      "test_cases/expectedOutput2.txt",
		"Hello\\nThere": "test_cases/expectedOutput3.txt",
		"{Hello There}": "test_cases/expectedOutput4.txt",
		"1Hello 2There": "test_cases/expectedOutput5.txt",
	}

	for input, expectedFilePath := range test_cases {
		//subtest to get the content from the file in the specified filepath
		t.Run(input, func(t *testing.T) {
			expectedContent, err := os.ReadFile(expectedFilePath)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", expectedFilePath, input, err)
			}

			//convert content read from the file to string
			expectedContentStr := string(expectedContent)

			result := functions.AsciiArt(input, fileLine)
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", input, expectedContentStr, result)
			}
		})
	}

}

func TestAsciiArtShadow(t *testing.T) {

	inputFile := "shadow.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var fileLine []string

	fileLine = strings.Split(string(file), "\n")

	//create a map containing the input text as key and the path to its expected output as value
	test_cases := map[string]string{
		"hello world":                "test_cases/expectedOutput6.txt",
		"123":                        "test_cases/expectedOutput7.txt",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ": "test_cases/expectedOutput8.txt",
	}

	for input, expectedFilePath := range test_cases {
		//subtest to get the content from the file in the specified filepath
		t.Run(input, func(t *testing.T) {
			expectedContent, err := os.ReadFile(expectedFilePath)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", expectedFilePath, input, err)
			}

			//convert content read from the file to string
			expectedContentStr := string(expectedContent)

			result := functions.AsciiArt(input, fileLine)
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", input, expectedContentStr, result)
			}
		})
	}

}

func TestAsciiArtThinkertoy(t *testing.T) {

	inputFile := "thinkertoy.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var fileLine []string

	fileLine = strings.Split(string(file), "\r\n")

	//create a map containing the input text as key and the path to its expected output as value
	test_cases := map[string]string{
		"nice 2 meet you": "test_cases/expectedOutput9.txt",
		"/(\")":           "test_cases/expectedOutput10.txt",
		"\"#$%&/()*+,-./": "test_cases/expectedOutput11.txt",
		"It's Working":    "test_cases/expectedOutput12.txt",
	}

	for input, expectedFilePath := range test_cases {
		//subtest to get the content from the file in the specified filepath
		t.Run(input, func(t *testing.T) {
			expectedContent, err := os.ReadFile(expectedFilePath)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", expectedFilePath, input, err)
			}

			//convert content read from the file to string
			expectedContentStr := string(expectedContent)

			result := functions.AsciiArt(input, fileLine)
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", input, expectedContentStr, result)
			}
		})
	}

}
