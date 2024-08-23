package utils_test

import (
	"bufio"
	"strings"
	"testing"

	"go.smsk.dev/todoman/utils"
)

func TestReadInput(t *testing.T) {
	// simulate user input
	input := "Hello, World!"
	// create a new reader with the input
	reader := bufio.NewReader(strings.NewReader(input + "\n"))
	// set the reader to the standard input
	utils.Reader = reader
	// call the ReadInput function
	result := utils.ReadInput()
	// check if the result is equal to the input
	// if not, fail the test
	if result != input {
		t.Fail()
	}
}

// Prompt function to prompt the user for input
func TestPrompt(t *testing.T) {
	// simulate user input
	input := "Hello, World!"
	// create a new reader with the input
	reader := bufio.NewReader(strings.NewReader(input + "\n"))
	// set the reader to the standard input
	utils.Reader = reader
	// call the Prompt function
	result := utils.Prompt("Enter something: ")
	// check if the result is equal to the input
	// if not, fail the test
	if result != input {
		t.Error("Expected:", input, "Got:", result)
	}
}
