package concurrency

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Output struct {
	Chart Chart `json:"chart"`
}

type Chart struct {
	Result []Result `json:"result"`
}

type Result struct {
	Meta Meta `json:"meta"`
}

type Meta struct {
	Currency           string  `json:"currency"`
	RegularMarketPrice float32 `json:"regularMarketPrice"`
	ChartPreviousClose float32 `json:"chartPreviousClose"`
	Error              error
}

func testConcurrentWorkers() {
	startDate := time.Now()
	average, err := monthlyData(startDate)

	fmt.Printf("This month the average value of the usd is: %f with error: %v\n", average, err)
}

func monthlyData(startDate time.Time) (float32, error) {
	daysInMonth := 0
	ch := make(chan Meta)
	date := startDate

	var averageValue float32

	for date.Month() == startDate.Month() {
		go getUSDValueForDay(date, ch)
		daysInMonth++
		date = date.Add(-24 * time.Hour)
	}

	for i := 0; i < daysInMonth; i++ {
		meta := <-ch
		if meta.Error != nil {
			fmt.Println(meta.Error)
			return 0, nil
		}
		averageValue += meta.RegularMarketPrice
	}

	averageValue /= float32(daysInMonth)

	return averageValue, nil
}

func getUSDValueForDay(date time.Time, ch chan<- Meta) {
	metaData := Meta{}
	defer func() {
		ch <- metaData
	}()

	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/COP=X?regular_market_time=%d&interval=1d", date.Unix())

	resp, err := http.Get(url)
	if err != nil {
		metaData.Error = err
		return
	}

	if resp.StatusCode != http.StatusOK {
		metaData.Error = errors.New("invalid return code")
	}

	defer resp.Body.Close()

	var responseData Output
	bodyData, err := io.ReadAll(resp.Body)

	if err != nil {
		metaData.Error = err
		return
	}

	err = json.Unmarshal(bodyData, &responseData)
	if err != nil {
		metaData.Error = err
		return
	}

	metaData = responseData.Chart.Result[0].Meta
}

// https://query1.finance.yahoo.com/v8/finance/chart/COP=X?date=02/10/2023&interval=1d
