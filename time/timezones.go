package time

import (
	"fmt"
	"time"
)

// This script will show how to work with timezones in go

func testTimeZones() {
	estTimeZone, _ := time.LoadLocation("EST") // first we load the desired timezones (o timezones in this case)
	chicagoTimeZone, _ := time.LoadLocation("America/Chicago")

	testTimeEST := time.Date(1994, 03, 28, 11, 20, 01, 00, estTimeZone) // after that we can load a date with a timezone
	testTimeChi := testTimeEST.In(chicagoTimeZone)                      // also we can convert a time to a different timezone

	fmt.Println("We can have a time in different timezones, EST: ", testTimeEST, " and also in chicago time: ", testTimeChi)
}
