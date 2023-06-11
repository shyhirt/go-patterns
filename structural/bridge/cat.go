package bridge

import "log"

type Cat struct{}

func (c *Cat) AnimalSound() {
	log.Println("Cat purr purr")
}
