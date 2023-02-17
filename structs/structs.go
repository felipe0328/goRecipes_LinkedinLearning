package structs

import (
	"fmt"
	"time"
)

type event struct {
	id          int
	description string
}

type light struct {
	event    event
	turnedOn bool
	hour     time.Time
}

type door struct {
	event  event
	locked bool
	hour   time.Time
}

const (
	lightEvent             = 1
	doorEvent              = 2
	doorOppenedDescription = "Door Opened"
	doorClosedDescription  = "Door Closed"
	lightTurnedOn          = "light On"
	lightTurnedOff         = "light Off"
)

func testStructs() {
	doorOpened := newDoorEvent(false)
	doorClosed := newDoorEvent(true)
	lightOn := newLightEvent(true)
	lightOff := newLightEvent(false)

	fmt.Printf("The events are door opened %v door closed %v light on %v light off %v\n", doorOpened, doorClosed, lightOn, lightOff)
}

func newDoorEvent(wasLocked bool) *door {
	doorEvent := &door{
		event:  event{id: doorEvent},
		locked: wasLocked,
		hour:   time.Now(),
	}

	doorEvent.event.description = doorOppenedDescription

	if wasLocked {
		doorEvent.event.description = doorClosedDescription
	}

	return doorEvent
}

func newLightEvent(turnedOn bool) *light {
	lightEvent := &light{
		event:    event{id: lightEvent},
		turnedOn: turnedOn,
		hour:     time.Now(),
	}

	lightEvent.event.description = lightTurnedOff

	if turnedOn {
		lightEvent.event.description = lightTurnedOn
	}

	return lightEvent
}
