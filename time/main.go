package time

import "fmt"

func TestTime() {
	fmt.Printf("---\n \tTime\n---\n")
	testArithmetic()
	testMeasureTime()
	testFormat()
	testParsing()
	testTimeZones()
	testParserAndConverter()
}
