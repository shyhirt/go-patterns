package concrete

type Bottle struct {
	Color       string
	Destination string
	Cap         float32
}

func (b *Bottle) SetCap(cap float32) {
	b.Cap = cap
}

func (b *Bottle) GetCap() float32 {
	return b.Cap
}

func (b *Bottle) SetColor(color string) {
	b.Color = color
}

func (b *Bottle) GetColor() string {
	return b.Color
}

func (b *Bottle) SetDestination(destination string) {
	b.Destination = destination
}

func (b *Bottle) GetDestination() string {
	return b.Destination
}
