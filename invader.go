package main

type Invader struct {
	pos    Point
	width  int8
	height int8
	value  int32
	isDead bool
}

func NewInvader(xPos int32, yPos int32) *Invader {
	return &Invader{pos: Point{}}
}
