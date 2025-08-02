package main

import "space-invaders/utils"

type SpriteUIComponent struct {
	sprite  string
	topLeft utils.Point
}

func (s SpriteUIComponent) getTopLeft() utils.Point {
	return s.topLeft
}

func (s SpriteUIComponent) getRasterized() string {
	return s.sprite
}

func NewSpriteUIComponent(sprite string, topLeft utils.Point) AbstractUiComponent {
	return SpriteUIComponent{sprite: sprite, topLeft: topLeft}
}
