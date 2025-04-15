package main

type SpriteUIComponent struct {
	sprite  string
	topLeft Point
}

func (s SpriteUIComponent) getTopLeft() Point {
	return s.topLeft
}

func (s SpriteUIComponent) getRasterized() string {
	return s.sprite
}

func NewSpriteUIComponent(sprite string, topLeft Point) AbstractUiComponent {
	return SpriteUIComponent{sprite: sprite, topLeft: topLeft}
}
