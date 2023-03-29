package point

import "fmt"

type Point struct {
	x, y int
}

func New(x, y int) Point {
	return Point{x: x, y: y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p Point) GetX() int {
	return p.x
}

func (p Point) GetY() int {
	return p.y
}

func (p Point) North() Point {
	return Point{
		x: p.x,
		y: p.y - 1,
	}
}
func (p Point) South() Point {
	return Point{
		x: p.x,
		y: p.y + 1,
	}
}

func (p Point) East() Point {
	return Point{
		x: p.x + 1,
		y: p.y,
	}
}

func (p Point) West() Point {
	return Point{
		x: p.x - 1,
		y: p.y,
	}
}
