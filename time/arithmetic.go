package time

import (
	"fmt"
	"time"
)

// This script will try to show the difference between time.Time and time.Duration
// and how we can work with both to do arithmetic time comparations

func testArithmetic() {
	testDate := time.Date(2023, 2, 16, 10, 0, 0, 0, time.UTC)

	isBDay := isBusinessDay(testDate)
	nextBusinessDay := getNextBusinessDay(testDate)
	nextWeekendDay := getNextWeekendDay(testDate)

	fmt.Println("for the date: ", testDate, " is businessDay: ", isBDay, " and the next BusinessDay is: ", nextBusinessDay, " and the next weekend day is: ", nextWeekendDay)
}

func isBusinessDay(date time.Time) bool {
	wDay := date.Weekday() // it converts the date to a weekday (Monday, Tuesday, ...)

	if wDay == time.Saturday || wDay == time.Sunday {
		return false
	}

	return true
}

func getNextBusinessDay(date time.Time) time.Time {
	day := time.Hour * 24 // day is a time.Duration value

	for {
		date = date.Add(day) // here we are using time.Duration to create new time.Time
		if isBusinessDay(date) {
			break
		}
	}

	return date
}

func getNextWeekendDay(date time.Time) time.Time {
	day := time.Hour * 24

	for {
		date = date.Add(day)
		if !isBusinessDay(date) {
			break
		}
	}

	return date
}
