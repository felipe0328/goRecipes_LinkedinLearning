package basics

import (
	"fmt"
	"regexp"
	"strings"
)

func testFrequency() {
	testString := "To be, or not to be. In this be world of world to."

	result := frequency(testString)

	fmt.Println("The frequency of the text: ", testString, "is ", result)
}

func frequency(input string) map[string]int {
	result := make(map[string]int)

	splitString := strings.Split(input, " ")

	regexExpression := regexp.MustCompile(`\W`)

	for _, data := range splitString {
		fixedData := strings.ToLower(data)
		fixedData = regexExpression.ReplaceAllString(fixedData, "")

		result[fixedData]++
	}

	return result
}
