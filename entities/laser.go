package entities

import (
	"space-invaders/ui"
	"space-invaders/utils"
)

type Laser struct {
	Position utils.Point
	dir      int // +/- 1
}

func (l *Laser) Update() {
	l.Position = l.Position.Add(utils.Point{X: 0, Y: l.dir})
}

func NewLaser(at *utils.Point, dir int) *Laser {
	return &Laser{
		Position: *at,
		dir:      dir,
	}
}

func (l *Laser) GetUI() []ui.AbstractUiComponent {
	return []ui.AbstractUiComponent{ui.NewSpriteUIComponent("║", l.Position)}
}
