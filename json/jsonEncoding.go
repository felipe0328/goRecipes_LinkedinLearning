package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// This sctruct will show how to work with json encoders

type Record struct {
	Time    time.Time
	Station string
	Temp    float64 `json:"temperature"`
	Rain    float64
}

func testJsonEncoding() {
	testingFile, err := os.Open("json/jsonValue.json")
	if err != nil {
		println(err)
		return
	}
	defer testingFile.Close()

	result, err := readRecord(testingFile)

	fmt.Printf("Decoding the entry returned %v with error %v\n", result, err)
}

func readRecord(r io.Reader) (*Record, error) {
	var record Record
	dec := json.NewDecoder(r)
	if err := dec.Decode(&record); err != nil {
		return nil, err
	}

	return &record, nil
}
