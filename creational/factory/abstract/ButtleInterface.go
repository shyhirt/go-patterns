package abstract

type BottleInterface interface {
	SetCap(cap float32)
	GetCap() float32
	SetColor(color string)
	GetColor() string
	SetDestination(destination string)
	GetDestination() string
}
