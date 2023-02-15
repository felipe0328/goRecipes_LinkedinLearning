package basics

import (
	"fmt"
	"regexp"
	"strings"
)

func TestFrequency() {
	testString := "To be, or not to be. In this be world of world to."

	result := frequency(testString)

	fmt.Println("The frecuency of the text: ", testString, "is ", result)
}

func frequency(input string) map[string]int {
	result := make(map[string]int)

	splittedString := strings.Split(input, " ")

	regexExpression := regexp.MustCompile(`\W`)

	for _, data := range splittedString {
		fixedData := strings.ToLower(data)
		fixedData = regexExpression.ReplaceAllString(fixedData, "")

		result[fixedData]++
	}

	return result
}
