package main

type Invader struct {
	pos    Point
	value  int32
	isDead bool
}

const INVADER_H = 1
const INVADER_W = 3

func NewInvader(xPos int, yPos int) *Invader {
	return &Invader{pos: Point{x: int(xPos), y: int(yPos)}, value: 1, isDead: false}
}

// move invader by specified distance
func (i *Invader) moveBy(x int, y int) {
	i.pos.x += x
	i.pos.y += y
}

func (i *Invader) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{NewSpriteUIComponent("v", i.topLeft())}
}

func (i *Invader) topLeft() Point {
	return i.pos
}
