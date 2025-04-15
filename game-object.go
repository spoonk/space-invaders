package main

type GameObject interface {
	moveTo(Point)
	topLeft() Point
	getUI() []AbstractUiComponent
}
