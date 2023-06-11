package builder

import (
	"log"
	"testing"
)

func Test_Builder(t *testing.T) {
	houseBuilder := getBuilder(HouseType)
	garageBuilder := getBuilder(GarageType)

	managerHouseBuilder := newBuildManager(houseBuilder)
	managerGarageBuilder := newBuildManager(garageBuilder)

	house := managerHouseBuilder.createBuild()
	garage := managerGarageBuilder.createBuild()

	log.Println("Garage characteristics:")
	log.Printf("Name: %s", garage.name)
	log.Printf("Height: %.2f", garage.height)
	log.Printf("Widgth: %.2f", garage.width)
	log.Printf("Lenght: %.2f", garage.length)

	log.Println("House characteristics:")
	log.Printf("Name: %s", house.name)
	log.Printf("Height: %.2f", house.height)
	log.Printf("Widgth: %.2f", house.width)
	log.Printf("Lenght: %.2f", house.length)
}
