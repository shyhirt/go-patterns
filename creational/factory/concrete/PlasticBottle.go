package concrete

import "github.com/shyhirt/go-patterns/creational/factory/abstract"

type PlasticBottle struct {
	Bottle
}

func NewPlasticBottle() abstract.BottleInterface {
	return &PlasticBottle{
		Bottle: Bottle{
			Color:       "transparent",
			Cap:         1,
			Destination: "Bottle for mineral water",
		},
	}
}
