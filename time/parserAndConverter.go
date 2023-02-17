package time

import (
	"fmt"
	"time"
)

const timeZoneFormat = "2006-01-02T03:04"

func testParserAndConverter() {
	testTime := "1994-03-28T11:00"
	resultTime := parserAndConverter(testTime, "EST", "Asia/Jerusalem")

	fmt.Println("Having the original time: ", testTime, " is this time in jerusalem: ", resultTime)

}

// Omiting handling with errors to keep this code as simple as possible
func parserAndConverter(inputTime, from, to string) time.Time {
	fromLocation, _ := time.LoadLocation(from)
	toLocation, _ := time.LoadLocation(to)

	originalTime, _ := time.ParseInLocation(timeZoneFormat, inputTime, fromLocation)

	return originalTime.In(toLocation)
}
