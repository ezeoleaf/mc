package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"log"
)

// Stack represents a character of type rune and the position of it.
type Stack struct {
	char  rune
	index int
}

var openingChars = map[rune]rune{
	rune('('): rune(')'),
	rune('{'): rune('}'),
	rune('['): rune(']'),
	rune('<'): rune('>'),
}

var closinChars = map[rune]rune{
	rune(')'): rune('('),
	rune('}'): rune('{'),
	rune(']'): rune('['),
	rune('>'): rune('<'),
}

func isValidChunk(chunks string) int {
	stack := []Stack{}

	for i, char := range chunks {
		// We store the opening chars in a stack.
		if _, found := openingChars[char]; found {
			stack = append(stack, Stack{
				char:  char,
				index: i,
			})
		} else if openingChar, found := closinChars[char]; found { // If it's not a opening char, we look it in the closing char list. If it's we retrieve the equivalent opening char.
			if len(stack) == 0 || openingChar != stack[len(stack)-1].char {
				return i
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return stack[0].index
	}

	// If stack is empty, it means that all opening chars have a correspondent closing char.
	return -1
}

func main() {
	filePath := envString("FILE_PATH", "./input.txt")

	readFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Only one line expected in the file. If not we only pick the first one.
	fileScanner.Scan()
	input := fileScanner.Text()

	result := isValidChunk(input)

	if result != -1 {
		fmt.Println(result)
		os.Exit(result)
	}
}

func envString(key string, fallback string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return fallback
}
