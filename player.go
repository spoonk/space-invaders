package main

type Player struct {
	pos Point
}

func (p *Player) moveTo(newPos Point) {
	p.pos = newPos
}

func (p *Player) topLeft() Point {
	return p.pos
}

func (p *Player) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{NewSpriteUIComponent("|||||\n|||||", p.topLeft())}
}

func NewPlayer() *Player {
	return &Player{pos: Point{50, 100}}
}
