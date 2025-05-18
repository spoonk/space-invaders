package main

type Invader struct {
	pos    Point
	value  int32
	isDead bool
}

// number of game tiles that an invader is tall & wide
const INVADER_W_H = 3

func NewInvader(xPos int, yPos int) *Invader {
	return &Invader{pos: Point{x: int(xPos), y: int(yPos)}, value: 1, isDead: false}
}

// move invader by specified distance
func (i *Invader) moveBy(x int, y int) {
	i.pos.x += x
	i.pos.y += y
}

func (i *Invader) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{NewSpriteUIComponent("-VV-", i.topLeft())}
}

func (i *Invader) topLeft() Point {
	return i.pos
}
