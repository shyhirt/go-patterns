package bridge

import "testing"

func TestCage(t *testing.T) {
	cat := &Cat{}
	dog := &Dog{}

	steelCage := &SteelCage{}
	glassCage := &GlassCage{}

	steelCage.SetAnimal(dog)
	steelCage.Sound()
	glassCage.SetAnimal(cat)
	glassCage.Sound()

	steelCage.SetAnimal(cat)
	steelCage.Sound()
	glassCage.SetAnimal(dog)
	glassCage.Sound()
}
