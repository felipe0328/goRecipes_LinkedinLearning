package time

import (
	"fmt"
	"net/http"
	"time"
)

// This script will show how to measure time taken by some code to execute

func testMeasureTime() {
	fmt.Println("The time it took to execute this method was: ", calculateTime(getSomeDatFromInternet))

}

func calculateTime(method func()) time.Duration {
	startTime := time.Now()
	method()
	return time.Since(startTime) // similar to time.Now().sub(startTime)
}

func getSomeDatFromInternet() {
	endpointURL := "https://google.com"

	_, _ = http.Get(endpointURL)
}
