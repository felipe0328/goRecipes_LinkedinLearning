package text

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

// This method will read a input value and then get how many times we used some go tool

var chllReg = regexp.MustCompile(`: (\d+):(\d);go ([a-zA-Z]+) .*`)

func testChallenge() {
	result, err := cmdFreq("text/challenge.txt")
	fmt.Printf("The challenge found the following result %v with err: %v\n", result, err)
}

func cmdFreq(fileName string) (map[string]int, error) {
	result := make(map[string]int)

	data, err := os.Open(fileName)
	if err != nil {
		return result, err
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		matches := chllReg.FindStringSubmatch(line)

		if matches == nil {
			continue
		}

		goFunc := matches[3]
		result[goFunc]++
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	if len(result) == 0 {
		return result, errors.New("no matches")
	}

	return result, nil
}
