package entities

import (
	"space-invaders/constants"
	"space-invaders/keyboard"
	"space-invaders/ui"
	"space-invaders/utils"
)

type Player struct {
	Pos        utils.Point
	controller *keyboard.KeyboardInputController
	Lives      int
}

func (p *Player) MoveTo(newPos utils.Point) {
	p.Pos = newPos
}

func (p *Player) TopLeft() utils.Point {
	return p.Pos
}

func (p *Player) GetUI() []ui.AbstractUiComponent {
	return []ui.AbstractUiComponent{ui.NewSpriteUIComponent("^-^", p.TopLeft())}
}

func NewPlayer() *Player {
	return &Player{Pos: utils.Point{X: constants.GAME_BOUNDARY.W / 2, Y: constants.PLAYER_Y}, controller: keyboard.GetController(), Lives: 3}
}

func (p *Player) RegisterHit() {
	p.Lives--
}

func (p *Player) BoundingBox() *utils.Box {
	return &utils.Box{X: p.Pos.X, Y: p.Pos.Y, H: 1, W: constants.PLAYER_W}
}

func (p *Player) Move() {
	press := p.controller.GetCurrentKeypress()

	currPos := p.Pos
	if press == 'a' {
		p.MoveTo(utils.Point{X: currPos.X - 1, Y: currPos.Y})
	} else if press == 'd' {
		p.MoveTo(utils.Point{X: currPos.X + 1, Y: currPos.Y})
	}
}
