package utils

import ()

type Box struct {
	X, Y, H, W int
}

func (b *Box) LeftBorderPos() int {
	return b.X
}

func (b *Box) RightBorderPos() int {
	return b.X + b.W
}

func (b *Box) GetTopLeft() *Point {
	return &Point{X: b.X, Y: b.Y}
}

func (b *Box) IsPointWithin(p *Point) bool {
	return (p.X >= b.X && p.X <= b.X+b.W && p.Y >= b.Y && p.Y <= b.Y+b.H)
}
