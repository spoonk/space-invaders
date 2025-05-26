package main

type Laser struct {
	position Point
	dir      int // +/- 1
}

func (l *Laser) update() {
	l.position = l.position.add(Point{x: 0, y: l.dir})
}

func NewLaser(at *Point, dir int) *Laser {
	return &Laser{
		position: *at,
		dir:      dir,
	}
}

func (l *Laser) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{NewSpriteUIComponent("║", l.position)}
}
