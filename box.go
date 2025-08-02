package main

import (
	"space-invaders/constants"
	"space-invaders/ui"
	"space-invaders/utils"
)

type Box struct {
	x, y, h, w int
}

func (b *Box) leftBorderPos() int {
	return b.x
}

func (b *Box) rightBorderPos() int {
	return b.x + b.w
}

func (b *Box) getTopLeft() *utils.Point {
	return &utils.Point{X: b.x, Y: b.y}
}

func (b *Box) isPointWithin(p *utils.Point) bool {
	return (p.X >= b.x && p.X <= b.x+b.w && p.Y >= b.y && p.Y <= b.y+b.h)
}

func (b *Box) getDebugUI() []ui.AbstractUiComponent {
	components := []ui.AbstractUiComponent{}

	if !constants.DEBUG_BOUNDARY {
		return components
	}

	// add in border characters
	for i := 1; i < b.w; i++ {
		components = append(components, ui.NewSpriteUIComponent("─", utils.Point{X: b.x + i, Y: b.y}))       // top
		components = append(components, ui.NewSpriteUIComponent("─", utils.Point{X: b.x + i, Y: b.y + b.h})) // bottom
	}

	for i := 1; i < b.h; i++ {
		components = append(components, ui.NewSpriteUIComponent("│", utils.Point{X: b.x, Y: b.y + i}))       // left
		components = append(components, ui.NewSpriteUIComponent("│", utils.Point{X: b.x + b.w, Y: b.y + i})) // right
	}

	components = append(components,
		[]ui.AbstractUiComponent{
			ui.NewSpriteUIComponent("╭", utils.Point{X: b.x, Y: b.y}),             // top left
			ui.NewSpriteUIComponent("╮", utils.Point{X: b.x + b.w, Y: b.y}),       // top right
			ui.NewSpriteUIComponent("╰", utils.Point{X: b.x, Y: b.y + b.h}),       // bot left
			ui.NewSpriteUIComponent("╯", utils.Point{X: b.x + b.w, Y: b.y + b.h}), // bot right
		}...,
	)

	return components
}
