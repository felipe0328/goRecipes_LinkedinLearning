package time

import (
	"fmt"
	"time"
)

// Parse time works to get a string time with some format and convert it to time.Time

func testParsing() {
	testTime := "28 Mar 1994"   // Time in string with specific format
	timeFormat := "02 Jan 2006" // Definition of format

	formattedTime, err := time.Parse(timeFormat, testTime)
	if err != nil {
		fmt.Println("The parse of the time was wrong, ", err)
	}

	// we can also format time durations
	durationTimeMs := "100ms"
	durationTimeNs := "10ns"
	durationTimeHour := "10h"
	composed := "10h5m10ns"

	ms, err := time.ParseDuration(durationTimeMs)
	if err != nil {
		fmt.Println("The parse of the time was wrong, ", err)
	}

	ns, err := time.ParseDuration(durationTimeNs)
	if err != nil {
		fmt.Println("The parse of the time was wrong, ", err)
	}

	h, err := time.ParseDuration(durationTimeHour)
	if err != nil {
		fmt.Println("The parse of the time was wrong, ", err)
	}

	c, err := time.ParseDuration(composed)
	if err != nil {
		fmt.Println("The parse of the time was wrong, ", err)
	}

	fmt.Println("We got the following results: Time: ", formattedTime, " millis: ", ms, " nanos: ", ns, " hours: ", h, " composed: ", c)
}
