package structs

import "fmt"

// this method will show how to work with ioat in go
// iota is similar to enum for other languages

type LogLevel uint8

const (
	DebugLevel LogLevel = iota + 1
	WarningLevel
	ErrorLevel
)

func testIota() {
	debugLevel := LogLevel(1)
	unrecognizedLevel := LogLevel(99)

	fmt.Printf("We defined two levels: %s and %s\n", debugLevel, unrecognizedLevel)
}

func (l LogLevel) String() string {
	var logLevel string
	switch l {
	case DebugLevel:
		logLevel = "Debug"
	case WarningLevel:
		logLevel = "Warning"
	case ErrorLevel:
		logLevel = "Error"
	default:
		logLevel = "Unrecognized"
	}

	return logLevel
}
