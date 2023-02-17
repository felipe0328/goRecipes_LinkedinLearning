package text

import (
	"fmt"
	"regexp"
	"strconv"
)

// this script shows how to work with regular expresions in golang
// so following the video example, I'm going to create a regular expresion for this type of entry
/*
	10 shares of MSFT for $123.4
	12 shares of TSLA for $567.8
*/

var reg = regexp.MustCompile(`(\d+) shares of ([A-Z]+) for \$(\d+(\.\d+)?)`) // this regular expresion matchs the entry

type transactionEntry struct {
	Symbol string
	Volume int
	Price  float64
}

func testRegex() {
	testData := []string{"15 shares of MSFT for $12.23", "1 shares of TWTR for $56.09"}

	results := make([]interface{}, len(testData))

	for index, value := range testData {
		transaction, err := parseLine(value)
		if err != nil {
			results[index] = err
			continue
		}

		results[index] = transaction
	}

	for index, values := range results {
		fmt.Printf("For this text %s, this was the result: %v ", testData[index], values)
	}

	fmt.Print("\n")
}

func parseLine(text string) (*transactionEntry, error) {
	matches := reg.FindStringSubmatch(text)
	if matches == nil {
		return nil, fmt.Errorf("no match in line: %s", text)
	}

	if len(matches) < 4 {
		return nil, fmt.Errorf("missing parameters in line: %s", text)
	}

	volume, vErr := strconv.Atoi(matches[1])
	symbol := matches[2]
	price, pErr := strconv.ParseFloat(matches[3], 64)

	if vErr != nil {
		return nil, fmt.Errorf("invalid int value: %s in text %s", matches[0], text)
	}

	if pErr != nil {
		return nil, fmt.Errorf("invalid float value %s in text %s", matches[2], text)
	}

	result := &transactionEntry{
		Volume: volume,
		Symbol: symbol,
		Price:  price,
	}

	return result, nil
}
