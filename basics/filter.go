package basics

import "fmt"

func testFilter() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := filter(isEven, input)
	resultOdd := filter(isOdd, input)

	fmt.Println("Having this list of numbers : ", input, " the following ones are even: ", result, " the following are odd: ", resultOdd)
}

func isEven(number int) bool {
	return number%2 != 0
}

func isOdd(number int) bool {
	return number%2 == 0
}

func filter(method func(int) bool, values []int) []int {
	var result []int

	for _, value := range values {
		if method(value) {
			result = append(result, value)
		}
	}

	return result
}
