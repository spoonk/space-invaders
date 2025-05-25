package main

type Invader struct {
	pos         Point
	boundingBox Box
	value       int32
	isDead      bool
}

const INVADER_H = 1
const INVADER_W = 3

func NewInvader(xPos int, yPos int) *Invader {
	return &Invader{
		pos:         Point{x: int(xPos), y: int(yPos)},
		value:       1,
		isDead:      false,
		boundingBox: Box{x: xPos, y: yPos, w: INVADER_W, h: INVADER_H}}
}

// move invader by specified distance
func (i *Invader) moveBy(x int, y int) {
	i.pos.x += x
	i.pos.y += y
	i.boundingBox.x += x
	i.boundingBox.y += y
}

func (i *Invader) getUI() []AbstractUiComponent {
	if i.isDead {
		return []AbstractUiComponent{}
	}

	return []AbstractUiComponent{NewSpriteUIComponent("▛▀▜", i.topLeft())}
}

func (i *Invader) topLeft() Point {
	return i.pos
}

func (i *Invader) registerHit() {
	i.isDead = true
	// todo increment score
}
