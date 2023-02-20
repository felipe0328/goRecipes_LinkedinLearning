package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// read rides.json and get maximun speed

type Ride struct {
	Start    string
	End      string
	Id       string
	Distance float64
}

func maxRideSpeed(r io.Reader) (float64, error) {
	dec := json.NewDecoder(r)

	var rideObject Ride

	var currentMaxSpeed float64

	for dec.Decode(&rideObject) == nil {
		startTime, _ := parseTime(rideObject.Start)
		endTime, _ := parseTime(rideObject.End)
		newSpeed := calculateSpeed(startTime, endTime, rideObject.Distance)
		if newSpeed > currentMaxSpeed {
			currentMaxSpeed = newSpeed
		}
	}

	return currentMaxSpeed, nil
}

func parseTime(input string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04", input)
}

func calculateSpeed(startTime, endTime time.Time, distance float64) float64 {
	dt := endTime.Sub(startTime)
	return distance / dt.Hours()
}

func testChallenge() {
	file, err := os.Open("json/rides.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	speed, err := maxRideSpeed(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(speed) // 40.5
}
