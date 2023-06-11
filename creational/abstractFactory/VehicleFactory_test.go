package abstractFactory

import (
	"log"
	"testing"
)

func TestVehicleFactory(t *testing.T) {
	bmw, err := VehicleFactory(Bmw)
	if err != nil {
		log.Fatalln(err)
	}
	bmwCar := bmw.MakeCar()
	log.Printf("Power: %f hp \r\n", bmwCar.GetPower())
	log.Printf("Has PC: %t  \r\n", bmwCar.HasComputer())
	log.Printf("Is all wheel drive: %t \r\n", bmwCar.IsAllWheelDrive())
	//Customize current BMW car
	bmwCar.SetAllWheelDrive(true)
	bmwCar.SetPower(1500)
	bmwCar.SetComputer(false)
	log.Printf("Power: %f hp \r\n", bmwCar.GetPower())
	log.Printf("Has PC: %t  \r\n", bmwCar.HasComputer())
	log.Printf("Is all wheel drive: %t \r\n", bmwCar.IsAllWheelDrive())

	//Create bike suzuki

	suzuki, err := VehicleFactory(Suzuki)
	if err != nil {
		log.Fatalln(err)
	}
	suzukiBike := suzuki.MakeBike()
	log.Printf("Power: %f hp \r\n", suzukiBike.GetPower())
	log.Printf("Has PC: %t  \r\n", suzukiBike.HasComputer())
	log.Printf("Color: %s \r\n", suzukiBike.GetColor())

	//Customize current Suzuki bike
	suzukiBike.SetColor("yellow")
	suzukiBike.SetPower(150)
	suzukiBike.SetComputer(false)

	log.Printf("Power: %f hp \r\n", suzukiBike.GetPower())
	log.Printf("Has PC: %t  \r\n", suzukiBike.HasComputer())
	log.Printf("Color: %s \r\n", suzukiBike.GetColor())

}
