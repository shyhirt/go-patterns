package abstract

type BikeInterface interface {
	SetColor(color string)
	GetColor() string
	SetComputer(hasComputer bool)
	HasComputer() bool
	SetPower(ePower float32) //Horsepower
	GetPower() float32
	SetWheels(wheels int)
	GetWheels() int
}
