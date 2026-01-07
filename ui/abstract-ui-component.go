package ui

import (
	"space-invaders/utils"
)

type AbstractUiComponent interface {
	GetTopLeft() utils.Point
	GetUI() []string
}

type StaticUI interface {
	GetTopLeft() utils.Point
	GetUI() []string
}
