package observer

import "log"

type EventObserver struct {
	ID any
}

func (e EventObserver) update(ev Event) {
	log.Printf("Send  %v to %v", ev, e.ID)
}
func (e EventObserver) getId() any {
	return e.ID
}
