package observer

type ObserverInterface interface {
	update(Event)
	getId() any
}
type Subject interface {
	attach(ObserverInterface)
	detach(ObserverInterface)
	notify(Event)
}

type Observer struct {
	ID any
}
