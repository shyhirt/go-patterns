package bridge

import "log"

type GlassCage struct {
	animal Animal
}

func (g *GlassCage) Sound() {
	log.Println("Sound from glass cage")
	g.animal.AnimalSound()
}

func (g *GlassCage) SetAnimal(animal Animal) {
	g.animal = animal
}
