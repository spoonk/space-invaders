package main

import (
	"space-invaders/constants"
	"space-invaders/utils"
)

type Player struct {
	pos        utils.Point
	controller *KeyboardInputController
	lives      int
}

func (p *Player) moveTo(newPos utils.Point) {
	p.pos = newPos
}

func (p *Player) topLeft() utils.Point {
	return p.pos
}

func (p *Player) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{NewSpriteUIComponent("^-^", p.topLeft())}
}

func NewPlayer() *Player {
	return &Player{pos: utils.Point{X: constants.GAME_BOUNDARY.W / 2, Y: constants.PLAYER_Y}, controller: GetController(), lives: 3}
}

func (p *Player) registerHit() {
	p.lives--
}

func (p *Player) boundingBox() *Box {
	return &Box{x: p.pos.X, y: p.pos.Y, h: 1, w: constants.PLAYER_W}
}

func (p *Player) move() {
	press := p.controller.getCurrentKeypress()

	currPos := p.pos
	if press == 'a' {
		p.moveTo(utils.Point{X: currPos.X - 1, Y: currPos.Y})
	} else if press == 'd' {
		p.moveTo(utils.Point{X: currPos.X + 1, Y: currPos.Y})
	}
}
