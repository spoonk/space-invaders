package main

type Player struct {
	pos        Point
	controller *KeyboardInputController
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
	return &Player{pos: Point{50, 31}, controller: GetController()}
}

func (p *Player) move() {
	press := p.controller.getCurrentKeypress()

	currPos := p.pos
	if press == 'a' {
		p.moveTo(Point{x: currPos.x - 1, y: currPos.y})
	} else if press == 'd' {
		p.moveTo(Point{x: currPos.x + 1, y: currPos.y})
	}
}
