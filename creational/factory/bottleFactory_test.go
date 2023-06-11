package factory

import (
	_ "github.com/shyhirt/go-patterns/creational/factory/concrete"
	"log"
	"testing"
)

func Test_getBottle(t *testing.T) {
	plasticBottle, err := getBottle(Plastic)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Default plastic bottle")
	log.Printf("Color: %s", plasticBottle.GetColor())
	log.Printf("Cap: %.2f liter", plasticBottle.GetCap())
	log.Printf("Destination: %s", plasticBottle.GetDestination())

	log.Println("Custom plastic bottle")
	plasticBottle.SetColor("black")
	plasticBottle.SetCap(2)
	plasticBottle.SetDestination("For juice")
	log.Printf("Color: %s", plasticBottle.GetColor())
	log.Printf("Cap: %.2f liter", plasticBottle.GetCap())
	log.Printf("Destination: %s", plasticBottle.GetDestination())

	glassBottle, err := getBottle(Glass)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Default glassBottle bottle")
	log.Printf("Color: %s", glassBottle.GetColor())
	log.Printf("Cap: %.2f liter", glassBottle.GetCap())
	log.Printf("Destination: %s", glassBottle.GetDestination())

	log.Println("Custom glass bottle")
	glassBottle.SetColor("black")
	glassBottle.SetCap(2)
	glassBottle.SetDestination("For juice")
	log.Printf("Color: %s", glassBottle.GetColor())
	log.Printf("Cap: %.2f liter", glassBottle.GetCap())
	log.Printf("Destination: %s", glassBottle.GetDestination())

}
