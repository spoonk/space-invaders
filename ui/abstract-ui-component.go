package ui

import "space-invaders/utils"

type AbstractUiComponent interface {
	GetTopLeft() utils.Point
	GetRasterized() string // bad, want to eventually do getSprite
}
