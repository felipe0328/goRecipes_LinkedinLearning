package time

import (
	"fmt"
	"time"
)

// Script showing different time.Format values

func testFormat() {
	testTime := time.Date(1994, 03, 28, 22, 10, 10, 0, time.FixedZone("COL", -500))

	fmt.Println("Here different formats to represent this time: ", testTime, " ---- ", testTime.Format(time.RFC822), " ---- ", testTime.Format("Mon, Jan 02"))
}
