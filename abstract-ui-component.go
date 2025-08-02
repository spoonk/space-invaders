package main

import "space-invaders/utils"

type AbstractUiComponent interface {
	getTopLeft() utils.Point
	getRasterized() string // bad, want to eventually do getSprite
}
