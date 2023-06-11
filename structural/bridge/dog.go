package bridge

import "log"

type Dog struct{}

func (d *Dog) AnimalSound() {
	log.Println("Dog woof woof")
}
