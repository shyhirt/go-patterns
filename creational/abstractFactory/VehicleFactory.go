package abstractFactory

import (
	"errors"
	"github.com/shyhirt/go-patterns/creational/abstractFactory/abstract"
	"github.com/shyhirt/go-patterns/creational/abstractFactory/concrete"
)

const (
	Bmw = iota
	Suzuki
)

var errVehicleFactory = errors.New("wrong model")

type VehicleInterface interface {
	MakeCar() abstract.CarInterface
	MakeBike() abstract.BikeInterface
}

func VehicleFactory(model int) (VehicleInterface, error) {
	switch model {
	case Bmw:
		return &concrete.BMW{}, nil
	case Suzuki:
		return &concrete.Suzuki{}, nil

	}

	return nil, errVehicleFactory
}
