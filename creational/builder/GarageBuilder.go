package builder

type GarageBuilder struct {
	length float32
	height float32
	width  float32
	name   string
}

func newGarageBuilder() *GarageBuilder {
	return &GarageBuilder{}
}

func (g *GarageBuilder) setWidth() {
	g.width = 100
}

func (g *GarageBuilder) setLength() {
	g.length = 70
}

func (g *GarageBuilder) setHeight() {
	g.height = 3
}

func (g *GarageBuilder) setName() {
	g.name = "Standard garage"
}

func (g *GarageBuilder) getBuild() Build {
	return Build{
		length: g.length,
		height: g.height,
		width:  g.width,
		name:   g.name,
	}
}
