package concrete

import "github.com/shyhirt/go-patterns/creational/abstractFactory/abstract"

type BMW struct {
}
type BMWCar struct {
	abstract.Car
}
type BMWBike struct {
	abstract.Bike
}

func (b *BMW) MakeCar() abstract.CarInterface {
	return &BMWCar{
		abstract.Car{
			Color:         "red",
			Computer:      true,
			Power:         1000,
			AllWheelDrive: true,
		},
	}
}

func (b *BMW) MakeBike() abstract.BikeInterface {
	return &BMWBike{abstract.Bike{
		Color:    "green",
		Computer: false,
		Power:    500,
		Wheels:   2,
	}}
}
