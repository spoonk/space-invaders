package main

import "space-invaders/utils"

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

func (l *Laser) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{NewSpriteUIComponent("║", l.position)}
}
