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
	return []AbstractUiComponent{NewSpriteUIComponent("^-^", p.topLeft())}
}

func NewPlayer() *Player {
	return &Player{pos: Point{50, 100}}
}

func (p *Player) move() {
	var press = kp
	if press == 97 {
		currPos := p.pos
		p.moveTo(Point{x: currPos.x - 1, y: currPos.y})
	} else if press == 100 {
		currPos := p.pos
		p.moveTo(Point{x: currPos.x + 1, y: currPos.y})
	}
}
