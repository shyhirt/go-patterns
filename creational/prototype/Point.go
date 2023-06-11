package prototype

type Point struct {
	X int
	Y int
}

func (p *Point) clone() PointInterface {
	return &Point{
		X: p.X,
		Y: p.Y,
	}
}

func (p *Point) incShuffleXY(x int, y int) {
	p.X = p.X + x
	p.Y = p.Y + y
}
func (p *Point) getPoint() *Point {
	return p
}
