package builder

const (
	HouseType = iota
	GarageType
)

type BuilderInterface interface {
	setHeight()
	setWidth()
	setLength()
	setName()
	getBuild() Build
}

func getBuilder(builderType int) BuilderInterface {
	switch builderType {
	case HouseType:
		return newHouseBuilder()
	case GarageType:
		return newGarageBuilder()
	}
	return nil
}
