package main

type Box struct {
	x, y, h, w int
}

func (b *Box) leftBorderPos() int {
	return b.x
}

func (b *Box) rightBorderPos() int {
	return b.x + b.w
}

func (b *Box) getTopLeft() *Point {
	return &Point{x: b.x, y: b.y}
}

func (b *Box) isPointWithin(p *Point) bool {
	return (p.x >= b.x && p.x <= b.x+b.w && p.y >= b.y && p.y <= b.y+b.h)
}

func (b *Box) getDebugUI() []AbstractUiComponent {
	components := []AbstractUiComponent{}

	if !DEBUG_BOUNDARY {
		return components
	}

	// add in border characters
	for i := 1; i < b.w; i++ {
		components = append(components, NewSpriteUIComponent("─", Point{x: b.x + i, y: b.y}))       // top
		components = append(components, NewSpriteUIComponent("─", Point{x: b.x + i, y: b.y + b.h})) // bottom
	}

	for i := 1; i < b.h; i++ {
		components = append(components, NewSpriteUIComponent("│", Point{x: b.x, y: b.y + i}))       // left
		components = append(components, NewSpriteUIComponent("│", Point{x: b.x + b.w, y: b.y + i})) // right
	}

	components = append(components,
		[]AbstractUiComponent{
			NewSpriteUIComponent("╭", Point{x: b.x, y: b.y}),             // top left
			NewSpriteUIComponent("╮", Point{x: b.x + b.w, y: b.y}),       // top right
			NewSpriteUIComponent("╰", Point{x: b.x, y: b.y + b.h}),       // bot left
			NewSpriteUIComponent("╯", Point{x: b.x + b.w, y: b.y + b.h}), // bot right
		}...,
	)

	return components
}
