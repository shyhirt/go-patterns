package prototype

import (
	"log"
	"testing"
)

func TestPoint_clone(t *testing.T) {
	point := Point{
		X: 40,
		Y: 50,
	}
	//Draw line Y
	var points []PointInterface
	for j := 0; j <= 100; j++ {
		tmpPoint := point.clone()
		tmpPoint.incShuffleXY(0, j)
		points = append(points, tmpPoint)
		log.Println(tmpPoint.getPoint())
	}
	//Line construction implementation from points array
	//....
	_ = points
}
