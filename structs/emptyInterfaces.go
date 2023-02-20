package structs

import "fmt"

// This script will show what is the purpose to work with empty interfaces
// The only interface should be used only as last resource
// the interface type is interface{} and is a variable type

var listOfData = make(map[string]int)

type DoorEvent struct{}

type TimerEvent struct{}

func testEmptyInterfaces() {
	doorEvent := &DoorEvent{}
	timerEvent := &TimerEvent{}
	doorEventNoPointer := DoorEvent{}

	addDataToList(doorEvent)
	addDataToList(timerEvent)
	addDataToList(doorEventNoPointer) // This is a no pointer object so we added an specific entry for that one

	fmt.Printf("We created multiple structs and added to the list, obtaining the following result: %v \n", listOfData)
}

func addDataToList(data interface{}) {
	fmt.Printf("%T\n", data)

	switch data.(type) {
	case *DoorEvent,
		DoorEvent:
		listOfData["DoorEvent"]++
	case *TimerEvent:
		listOfData["TimerEvent"]++
	default:
		break
	}
}
