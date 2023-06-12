package observer

import (
	"log"
	"testing"
)

func TestEventNotifier_remove(t *testing.T) {
	type fields struct {
		observers []ObserverInterface
	}
	type args struct {
		toBeRemoved ObserverInterface
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "test", fields: struct{ observers []ObserverInterface }{observers: []ObserverInterface{
			ObserverInterface(EventObserver{ID: 1}),
			ObserverInterface(EventObserver{ID: 2}),
			ObserverInterface(EventObserver{ID: 3}),
			ObserverInterface(EventObserver{ID: 4}),
		}}, args: args{toBeRemoved: ObserverInterface(EventObserver{ID: 2})},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventNotifier{
				observers: tt.fields.observers,
			}
			e.remove(tt.args.toBeRemoved)
			log.Println(len(e.observers), e.observers)
		})
	}
}
