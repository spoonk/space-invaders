package main

type Invader struct {
	topLeft Point
	width   int8
	height  int8
	value   int32
	isDead  bool
}

// number of game tiles that an invader is tall & wide
const INVADER_W_H = 3

func NewInvader(xPos int, yPos int) *Invader {
	return &Invader{topLeft: Point{x: int(xPos), y: int(yPos)}}
}

// move invader by specified distance
func (i *Invader) moveBy(x int, y int) {
	i.topLeft.x += x
	i.topLeft.y += y
}

func (i *Invader) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{NewSpriteUIComponent("-VV-", i.topLeft)}
}
