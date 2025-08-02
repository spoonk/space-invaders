package main

import (
	"space-invaders/ui"
	"space-invaders/utils"
)

type Laser struct {
	position utils.Point
	dir      int // +/- 1
}

func (l *Laser) update() {
	l.position = l.position.Add(utils.Point{X: 0, Y: l.dir})
}

func NewLaser(at *utils.Point, dir int) *Laser {
	return &Laser{
		position: *at,
		dir:      dir,
	}
}

func (l *Laser) getUI() []ui.AbstractUiComponent {
	return []ui.AbstractUiComponent{ui.NewSpriteUIComponent("║", l.position)}
}
