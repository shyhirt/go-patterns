package builder

type HouseBuilder struct {
	length float32
	height float32
	width  float32
	name   string
}

func newHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (h *HouseBuilder) setWidth() {
	h.width = 200
}

func (h *HouseBuilder) setLength() {
	h.length = 140
}

func (h *HouseBuilder) setHeight() {
	h.height = 5
}

func (h *HouseBuilder) setName() {
	h.name = "Standard house"
}

func (h *HouseBuilder) getBuild() Build {
	return Build{
		length: h.length,
		height: h.height,
		width:  h.width,
		name:   h.name,
	}
}
