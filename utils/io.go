package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var Reader *bufio.Reader

// ReadInput function to read input from the user
func ReadInput() string {
	reader := bufio.NewReader(Reader)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Prompt function to prompt the user for input
func Prompt(prompt string) string {
	if prompt != "" {
		fmt.Print(prompt)
	}
	reader := bufio.NewReader(Reader)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// !deprecated! use gorm to handle database operations
// WriteToFile function to write content to a file
func WriteToFile(filename string, content string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
