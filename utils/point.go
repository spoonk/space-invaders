package utils

type Point struct {
	X, Y int
}

func (p *Point) Shifted(x int, y int) Point {
	return Point{X: p.X + x, Y: p.Y + y}
}

func (p *Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}
