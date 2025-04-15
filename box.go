package main

type Box struct {
	x, y, h, w int
}

func (b Box) leftBorderPos() int {
	return b.x
}

func (b Box) rightBorderPos() int {
	return b.x + b.w
}
