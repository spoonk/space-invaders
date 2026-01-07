package ui

import (
	"space-invaders/constants"
	"space-invaders/utils"
	"unicode/utf8"
)

type SpriteUIComponent struct {
	sprite  []string
	topLeft utils.Point
}

func (s SpriteUIComponent) GetTopLeft() utils.Point {
	return s.topLeft
}

func (s SpriteUIComponent) GetUI() []string {
	return s.sprite
}

func NewSpriteUIComponent(sprite string, topLeft utils.Point) AbstractUiComponent {
	return SpriteUIComponent{sprite: []string{sprite}, topLeft: topLeft}
}

func NewCenteredTextUIComponent(text string) AbstractUiComponent {
	length := utf8.RuneCountInString(text)
	centeredPosition := utils.Point{X: constants.GAME_BOUNDARY.W/2 - length/3, Y: constants.GAME_BOUNDARY.H / 2}
	return SpriteUIComponent{sprite: []string{text}, topLeft: centeredPosition}
}

func NewMultiLineSpriteUIComponent(sprite []string, topLeft utils.Point) AbstractUiComponent {
	return SpriteUIComponent{sprite: sprite, topLeft: topLeft}
}
