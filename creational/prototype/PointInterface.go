package prototype

type PointInterface interface {
	clone() PointInterface
	getPoint() *Point
	incShuffleXY(x int, y int)
}
