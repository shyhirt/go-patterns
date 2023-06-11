package concrete

import "github.com/shyhirt/go-patterns/creational/factory/abstract"

type GlassBottle struct {
	Bottle
}

func NewGlassBottle() abstract.BottleInterface {
	return &GlassBottle{
		Bottle: Bottle{
			Color:       "green",
			Cap:         1,
			Destination: "Bottle for lemonade",
		},
	}
}
