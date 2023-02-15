package basics

import "fmt"

func TestMean() {
	listOfNumbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Mean of numbers: ", listOfNumbers, " is: ", mean(listOfNumbers))
}

// mean is the sum of all the numbers divided by the amount of numbers
func mean(numbers []int) float64 {
	var result float64

	for _, number := range numbers {
		result += float64(number)
	}

	result /= float64(len(numbers))

	return result
}
