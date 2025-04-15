package main

// container of all aliens
type Wave struct {
	boundingBox Box
}

func (w Wave) moveTo() {
}

func (w Wave) leftBorderPos() int {
	return w.boundingBox.x
}

func (w Wave) rightBorderPos() int {
	return w.boundingBox.x + w.boundingBox.w
}

func (w Wave) topLeft() Point {
	return Point{x: w.boundingBox.x, y: w.boundingBox.y}
}

func (w Wave) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{
		NewSpriteUIComponent("╭", Point{x: w.boundingBox.x, y: w.boundingBox.y}),                                     // top left
		NewSpriteUIComponent("╮", Point{x: w.boundingBox.x + w.boundingBox.w, y: w.boundingBox.y}),                   // top right
		NewSpriteUIComponent("╰", Point{x: w.boundingBox.x, y: w.boundingBox.y + w.boundingBox.h}),                   // bot left
		NewSpriteUIComponent("╯", Point{x: w.boundingBox.x + w.boundingBox.w, y: w.boundingBox.y + w.boundingBox.h}), // bot right
	}
}
