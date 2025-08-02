package main

import (
	"space-invaders/constants"
	"space-invaders/ui"
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

func (p *Player) getUI() []ui.AbstractUiComponent {
	return []ui.AbstractUiComponent{ui.NewSpriteUIComponent("^-^", p.topLeft())}
}

func NewPlayer() *Player {
	return &Player{pos: utils.Point{X: constants.GAME_BOUNDARY.W / 2, Y: constants.PLAYER_Y}, controller: GetController(), lives: 3}
}

func (p *Player) registerHit() {
	p.lives--
}

func (p *Player) boundingBox() *utils.Box {
	return &utils.Box{X: p.pos.X, Y: p.pos.Y, H: 1, W: constants.PLAYER_W}
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
