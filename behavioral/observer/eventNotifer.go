package observer

type EventNotifier struct {
	observers []ObserverInterface
	data      any
}

func NewEventNotifier(data any) *EventNotifier {
	return &EventNotifier{
		data: data,
	}
}
func (e *EventNotifier) attach(o ObserverInterface) {
	e.observers = append(e.observers, o)
}

func (e *EventNotifier) detach(o ObserverInterface) {
	e.remove(o)
}
func (e *EventNotifier) remove(toBeRemoved ObserverInterface) {
	l := len(e.observers)
	for i, observer := range e.observers {
		if toBeRemoved == observer {
			e.observers[l-1], e.observers[i] = e.observers[i], e.observers[l-1]
			e.observers = e.observers[:l-1]
			return
		}
	}
}
func (e *EventNotifier) notify() {
	for _, observer := range e.observers {
		observer.update(Event{
			Data: e.data,
		})
	}
}
