package ui

import "space-invaders/utils"

type AbstractUiComponent interface {
	GetTopLeft() utils.Point
	GetRasterized() []string
}

type StaticUI interface {
	GetTopLeft() utils.Point
	GetRasterized() []string
}
