package text

import (
	"fmt"
	"unicode/utf8"
)

// This script will try to show how to work with unicodes in golang

func testUnicode() {
	testWords := []string{"Hi,", "I", "am", "your", "father"}
	length := entryLength(testWords)
	fmt.Printf("The amount of letters in the entry %v was %d\n", testWords, length)
}

func entryLength(words []string) int {
	var total int

	for _, value := range words {
		total += utf8.RuneCountInString(value) // this gets the length of the string, it is receiving a string but can get an array of bytes
	}

	return total
}
