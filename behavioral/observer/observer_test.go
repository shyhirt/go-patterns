package observer

import (
	"log"
	"testing"
)

func TestObserver(t *testing.T) {
	event := NewEventNotifier("Event text")

	observerOne := &EventObserver{ID: "ID_1"}
	observerTwo := &EventObserver{ID: "ID_2"}
	observerThree := &EventObserver{ID: "ID_3"}

	event.attach(observerOne)
	event.attach(observerTwo)
	event.attach(observerThree)
	event.notify()
	log.Println("")
	event.remove(observerOne)
	event.notify()
}
