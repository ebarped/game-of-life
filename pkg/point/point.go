package point

import "fmt"

type Point struct {
	X, Y int
}

func New(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Point) GetNorth() Point {
	return Point{
		X: p.X,
		Y: p.Y - 1,
	}
}
func (p Point) GetSouth() Point {
	return Point{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p Point) GetEast() Point {
	return Point{
		X: p.X + 1,
		Y: p.Y,
	}
}

func (p Point) GetWest() Point {
	return Point{
		X: p.X - 1,
		Y: p.Y,
	}
}
