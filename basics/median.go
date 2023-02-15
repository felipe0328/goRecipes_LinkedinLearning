package basics

import (
	"fmt"
	"sort"
)

func TestMedian() {
	numbers := []float64{1, 2, 3, 4, 5, 6}
	sortedResult := median(numbers)

	notSortedNumbers := []float64{5, 2, 3, 1, 4}
	notSortedResult := median(notSortedNumbers)

	fmt.Println("Median of numbers ", numbers, " is: ", sortedResult, " - (notSortedTest): ", notSortedResult)
}

// the median finds the value in the middle of the organized list of numbers passed
func median(numbers []float64) float64 {
	var result float64
	numbersCopy := make([]float64, len(numbers))
	copy(numbersCopy, numbers)

	if len(numbers) == 0 {
		return result
	}

	if len(numbers) == 1 {
		return float64(numbers[0])
	}

	// first we sort the input list
	sort.Float64s(numbersCopy)

	halfLength := len(numbers) / 2
	if halfLength%2 != 0 {
		result = (numbersCopy[halfLength-1] + numbersCopy[halfLength]) / 2
	} else {
		result = numbersCopy[halfLength]
	}

	return result
}
