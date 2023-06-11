package abstract

type Car struct {
	Color         string
	Computer      bool
	Power         float32
	AllWheelDrive bool
}

func (c *Car) SetComputer(hasComputer bool) {
	c.Computer = hasComputer
}

func (c *Car) HasComputer() bool {
	return c.Computer
}

func (c *Car) SetPower(ePower float32) {
	c.Power = ePower
}

func (c *Car) GetPower() float32 {
	return c.Power
}

func (c *Car) SetAllWheelDrive(allWheelDrive bool) {
	c.AllWheelDrive = allWheelDrive
}

func (c *Car) IsAllWheelDrive() bool {
	return c.AllWheelDrive
}
