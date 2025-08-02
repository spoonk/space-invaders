package main

import "space-invaders/utils"

type GameObject interface {
	moveTo(utils.Point)
	topLeft() utils.Point
	getUI() []AbstractUiComponent
}
