package main

import "space-invaders/constants"

type Player struct {
	pos        Point
	controller *KeyboardInputController
	lives      int
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
	return &Player{pos: Point{constants.GAME_BOUNDARY.W / 2, constants.PLAYER_Y}, controller: GetController(), lives: 3}
}

func (p *Player) registerHit() {
	p.lives--
}

func (p *Player) boundingBox() *Box {
	return &Box{x: p.pos.x, y: p.pos.y, h: 1, w: constants.PLAYER_W}
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
