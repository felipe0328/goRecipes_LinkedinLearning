package basics

import (
	"errors"
	"fmt"
)

var (
	errorEmpty    = errors.New("the input data is empty")
	errorNotEmpty = errors.New("the input is not empty")
)

// This is a simple test of an error handling
func TestError() {
	fmt.Println("Test Error if empty ", errorCreator(""), " if not empty ", errorCreator("something"), " and nil when ok: ", errorCreator("safeWord"))
}

func errorCreator(input string) error {
	if input == "" {
		return errorEmpty
	}

	if input == "safeWord" {
		return nil
	}

	return errorNotEmpty
}
