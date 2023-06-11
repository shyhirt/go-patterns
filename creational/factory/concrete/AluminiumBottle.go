package concrete

import "github.com/shyhirt/go-patterns/creational/factory/abstract"

type AluminiumBottle struct {
	Bottle
}

func NewAluminiumBottle() abstract.BottleInterface {
	return &AluminiumBottle{
		Bottle: Bottle{
			Color:       "silver",
			Cap:         0.5,
			Destination: "Bottle for beer",
		},
	}
}
