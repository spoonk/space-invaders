package ui

import "space-invaders/utils"

type SpriteUIComponent struct {
	sprite  []string
	topLeft utils.Point
}

func (s SpriteUIComponent) GetTopLeft() utils.Point {
	return s.topLeft
}

func (s SpriteUIComponent) GetRasterized() []string {
	return s.sprite
}

func NewSpriteUIComponent(sprite string, topLeft utils.Point) AbstractUiComponent {
	return SpriteUIComponent{sprite: []string{sprite}, topLeft: topLeft}
}

func NewMultiLineSpriteUIComponent(sprite []string, topLeft utils.Point) AbstractUiComponent {
	return SpriteUIComponent{sprite: sprite, topLeft: topLeft}
}
