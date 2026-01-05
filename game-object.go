package main

import (
	"space-invaders/ui"
	"space-invaders/utils"
)

type GameObject interface {
	moveTo(utils.Point)
	topLeft() utils.Point
	GetUI() []ui.StaticUI
	Container() utils.Box
}
