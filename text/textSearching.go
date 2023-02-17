package text

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// for this script we are going to show how to do a grep using go for a input text

func testSearching() {
	sampleText := createSampleText()
	wordToSearch := "the"

	result := searchForValues(sampleText, wordToSearch)

	fmt.Printf("In the test text we found the word 'the' in %d lines\n", len(result))
}

func searchForValues(reader io.Reader, param string) []string {
	var matches []string

	search := bufio.NewScanner(reader)

	for search.Scan() {
		textValue := search.Text()
		if strings.Contains(textValue, param) {
			matches = append(matches, textValue)
		}
	}

	return matches
}

func createSampleText() io.Reader {
	newBuf := new(bytes.Buffer)

	testEntry := []string{
		"This is the test string",
		"we are going to search",
		"the lines containing",
		"the word 'the'",
	}

	for _, value := range testEntry {
		newBuf.WriteString(fmt.Sprintf("%s\n", value))
	}

	return newBuf
}
