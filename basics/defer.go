package basics

import (
	"encoding/json"
	"fmt"
	"os"
)

// testDefer Simple test to check the defer method
func testDefer() {
	filename := "testFile.txt"
	createdError := createFile(filename, "testData")
	deletedError := deleteFile(filename)
	fmt.Println("Creating file ", filename, " with error: ", createdError, " and deleted with error: ", deletedError)
}

func createFile(filename string, data interface{}) (err error) {
	newFile, err := os.Create(filename)

	if err != nil {
		return
	}

	// using defer to close the file after finishing with the script
	defer newFile.Close()

	dataBytes, err := json.Marshal(data)

	if err != nil {
		return
	}

	defer newFile.Sync()
	newFile.Write(dataBytes)

	return
}

func deleteFile(filename string) error {
	// deleting the file after finishing with it
	err := os.Remove(filename)
	return err
}
