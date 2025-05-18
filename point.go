package main

type Point struct {
	x, y int
}

func (p *Point) shifted(x int, y int) Point {
	return Point{x: p.x + x, y: p.y + y}
}
