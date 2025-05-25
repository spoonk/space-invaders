package main

type Laser struct {
	position    Point
	boundingBox Box
}

func (l *Laser) update() {
	l.position = l.position.add(Point{x: 0, y: -1})
}

func NewLaser(at *Point) *Laser {
	return &Laser{
		position: *at,
		boundingBox: Box{
			x: at.x, y: at.y, h: 1, w: 1,
		},
	}
}

func (l *Laser) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{NewSpriteUIComponent("║", l.position)}
}
