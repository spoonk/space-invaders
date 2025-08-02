package ui

import (
	"space-invaders/constants"
	"space-invaders/utils"
)

// func (b *Box) getDebugUI() []ui.AbstractUiComponent {
func GetDebugBoxUI(b *utils.Box) []AbstractUiComponent {
	components := []AbstractUiComponent{}

	if !constants.DEBUG_BOUNDARY {
		return components
	}

	// add in border characters
	for i := 1; i < b.W; i++ {
		components = append(components, NewSpriteUIComponent("─", utils.Point{X: b.X + i, Y: b.Y}))       // top
		components = append(components, NewSpriteUIComponent("─", utils.Point{X: b.X + i, Y: b.Y + b.H})) // bottom
	}

	for i := 1; i < b.H; i++ {
		components = append(components, NewSpriteUIComponent("│", utils.Point{X: b.X, Y: b.Y + i}))       // left
		components = append(components, NewSpriteUIComponent("│", utils.Point{X: b.X + b.W, Y: b.Y + i})) // right
	}

	components = append(components,
		[]AbstractUiComponent{
			NewSpriteUIComponent("╭", utils.Point{X: b.X, Y: b.Y}),             // top left
			NewSpriteUIComponent("╮", utils.Point{X: b.X + b.W, Y: b.Y}),       // top right
			NewSpriteUIComponent("╰", utils.Point{X: b.X, Y: b.Y + b.H}),       // bot left
			NewSpriteUIComponent("╯", utils.Point{X: b.X + b.W, Y: b.Y + b.H}), // bot right
		}...,
	)

	return components
}
