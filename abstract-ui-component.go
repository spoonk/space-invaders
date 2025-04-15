package main

type AbstractUiComponent interface {
	getTopLeft() Point
	getRasterized() string // bad, want to eventually do getSprite
}

type Point struct {
	x, y int
}
