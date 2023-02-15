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
	emptyError := errorCreator("")
	notEmptyError := errorCreator("something")
	nilError := errorCreator("safeWord")

	fmt.Println("Test Error if empty: ", errorHandler(emptyError), " - if not empty: ", errorHandler(notEmptyError), " - and nil when ok: ", errorHandler(nilError))
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

func errorHandler(err error) string {
	switch err {
	case errorEmpty:
		return "handling if empty"
	case errorNotEmpty:
		return "handling when not empty"
	default:
		return "everything is ok"
	}
}
