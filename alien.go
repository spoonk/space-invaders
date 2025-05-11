package main

type Alien struct {
	pos    Point
	width  int8
	height int8
}

func NewAlien(xPos int32, yPos int32) *Alien {
	return &Alien{pos: Point{}}
}
