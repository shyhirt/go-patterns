package factory

import (
	"errors"
	"github.com/shyhirt/go-patterns/creational/factory/abstract"
	"github.com/shyhirt/go-patterns/creational/factory/concrete"
)

const (
	Plastic = iota
	Glass
	Aluminium
)

var errGetBottle = errors.New("error get bottle type")

func getBottle(bottleType int) (abstract.BottleInterface, error) {
	switch bottleType {
	case Plastic:
		return concrete.NewPlasticBottle(), nil
	case Glass:
		return concrete.NewGlassBottle(), nil
	case Aluminium:
		return concrete.NewAluminiumBottle(), nil
	}
	return nil, errGetBottle
}
