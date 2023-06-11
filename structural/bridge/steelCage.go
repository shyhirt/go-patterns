package bridge

import "log"

type SteelCage struct {
	animal Animal
}

func (s *SteelCage) Sound() {
	log.Println("Sound from steel cage")
	s.animal.AnimalSound()
}

func (s *SteelCage) SetAnimal(animal Animal) {
	s.animal = animal
}
