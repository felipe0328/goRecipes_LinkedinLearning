package text

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func testGenerateReport() {
	indexValue := 1
	something := "something"
	log.Printf("testing %d: %#v", indexValue, something) // the log package logs with a timestamp

	buf := new(bytes.Buffer) // the writer can come from different sources, here I'm using a bytes buffer

	testingMessages := []string{"This is a output message,", "Formed by different messages."}
	generateReport(buf, testingMessages)
	fmt.Println(buf.String())
}

func generateReport(w io.Writer, messages []string) {
	for index, value := range messages {
		_, err := fmt.Fprintf(w, "%-10s \tindex:%04d\n", value, index)
		if err != nil {
			return
		}
	}
}
