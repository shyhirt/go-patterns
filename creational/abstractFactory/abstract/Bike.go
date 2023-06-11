package abstract

type Bike struct {
	Color    string
	Computer bool
	Power    float32
	Wheels   int
}

func (b *Bike) SetColor(color string) {
	b.Color = color
}

func (b *Bike) GetColor() string {
	return b.Color
}

func (b *Bike) SetComputer(hasComputer bool) {
	b.Computer = hasComputer
}

func (b *Bike) HasComputer() bool {
	return b.Computer
}

func (b *Bike) SetPower(ePower float32) {
	b.Power = ePower
}

func (b *Bike) GetPower() float32 {
	return b.Power
}

func (b *Bike) SetWheels(wheels int) {
	b.Wheels = wheels
}

func (b *Bike) GetWheels() int {
	return b.Wheels
}
