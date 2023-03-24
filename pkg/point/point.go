package point

type Point struct {
	x, y int
}

func New(x, y int) Point {
	return Point{x: x, y: y}
}
