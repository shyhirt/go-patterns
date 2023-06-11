package abstract

type CarInterface interface {
	SetComputer(hasComputer bool)
	HasComputer() bool
	SetPower(ePower float32) //Horsepower
	GetPower() float32
	SetAllWheelDrive(allWheelDrive bool)
	IsAllWheelDrive() bool
}
