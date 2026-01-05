package entities

import (
	"space-invaders/ui"
	"space-invaders/utils"
)

type Invader struct {
	pos         utils.Point
	BoundingBox utils.Box
	Value       int32
	IsDead      bool
}

const (
	INVADER_H = 1
	INVADER_W = 3
)

func NewInvader(xPos int, yPos int) *Invader {
	return &Invader{
		pos:         utils.Point{X: int(xPos), Y: int(yPos)},
		Value:       1,
		IsDead:      false,
		BoundingBox: utils.Box{X: xPos, Y: yPos, W: INVADER_W, H: INVADER_H},
	}
}

// move invader by specified distance
func (i *Invader) moveBy(x int, y int) {
	i.pos.X += x
	i.pos.Y += y
	i.BoundingBox.X += x
	i.BoundingBox.Y += y
}

func (i *Invader) GetStaticUI() []ui.StaticUI {
	if i.IsDead {
		return []ui.StaticUI{}
	}

	return append(
		[]ui.StaticUI{ui.NewSpriteUIComponent("▛▀▜", i.topLeft())},
		ui.GetDebugBoxUI(&i.BoundingBox)...,
	)
}

var image [][]float64 = nil

func (i *Invader) GetImage() *[][]float64 {
	if image == nil {
		image, _ = utils.ReadImageToFloat64("invader.png")
	}
	return &image
}

func (i *Invader) topLeft() utils.Point {
	return i.pos
}

func (i *Invader) RegisterHit() {
	i.IsDead = true
}
