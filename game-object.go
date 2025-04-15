package main

type GameObject interface {
	moveTo()
	topLeft() Point
	getUI() []AbstractUiComponent
}
