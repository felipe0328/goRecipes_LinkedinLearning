package basics

import "fmt"

func testPanic() {
	inputArray := []int{1}
	fmt.Println("The recover returned the folowing error: ", runMethodSafe(getArrayLenMinusTwoValue, inputArray))
	// getArrayLenMinusTwoValue(inputArray) This will fail and will stop the process
}

func runMethodSafe(method func(params []int) int, input interface{}) (result interface{}) {
	defer func() {
		if err := recover(); err != nil {
			result = err
		}
	}()

	return method(input.([]int))
}

func getArrayLenMinusTwoValue(input []int) int {
	return input[len(input)-2]
}
