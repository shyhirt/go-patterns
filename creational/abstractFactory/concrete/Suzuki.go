package concrete

import (
	"github.com/shyhirt/go-patterns/creational/abstractFactory/abstract"
)

type Suzuki struct {
}
type SuzukiCar struct {
	abstract.Car
}
type SuzukiBike struct {
	abstract.Bike
}

func (b *Suzuki) MakeCar() abstract.CarInterface {
	return &SuzukiCar{
		abstract.Car{
			Color:         "red",
			Computer:      true,
			Power:         1000,
			AllWheelDrive: true,
		},
	}
}

func (b *Suzuki) MakeBike() abstract.BikeInterface {
	return &SuzukiBike{abstract.Bike{
		Color:    "green",
		Computer: false,
		Power:    500,
		Wheels:   2,
	}}
}
