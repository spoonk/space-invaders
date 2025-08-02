package main

import (
	"space-invaders/ui"
	"space-invaders/utils"
)

type Invader struct {
	pos         utils.Point
	boundingBox Box
	value       int32
	isDead      bool
}

const (
	INVADER_H = 1
	INVADER_W = 3
)

func NewInvader(xPos int, yPos int) *Invader {
	return &Invader{
		pos:         utils.Point{X: int(xPos), Y: int(yPos)},
		value:       1,
		isDead:      false,
		boundingBox: Box{x: xPos, y: yPos, w: INVADER_W, h: INVADER_H},
	}
}

// move invader by specified distance
func (i *Invader) moveBy(x int, y int) {
	i.pos.X += x
	i.pos.Y += y
	i.boundingBox.x += x
	i.boundingBox.y += y
}

func (i *Invader) getUI() []ui.AbstractUiComponent {
	if i.isDead {
		return []ui.AbstractUiComponent{}
	}

	return append(
		[]ui.AbstractUiComponent{ui.NewSpriteUIComponent("▛▀▜", i.topLeft())},
		i.boundingBox.getDebugUI()...,
	)
}

func (i *Invader) topLeft() utils.Point {
	return i.pos
}

func (i *Invader) registerHit() {
	i.isDead = true
}
