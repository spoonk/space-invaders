package main

// wave designed as a collection of invaders
// wave is responsible for moving all invaders within

type InvaderWave struct {
	boundingBox  Box
	gameBoundary *Box
	invaders     [][]*Invader
	currentDir   string
}

// TODO: different types of invaders

const Y_SPEED = 1
const WAVE_HEIGHT = 5
const WAVE_WIDTH = 11

// wave needs to define a delta T for how much it should move..
// should be hitching independent: project forward by dt, which can
// lead to you moving > 1 position in a single update

// each wave has the exact same config of invaders
func NewInvaderWave(gameBoundary *Box) *InvaderWave {

	// this process is backwards, the bounding box should be inferred from the set of invaders
	waveBoundingBox := Box{
		x: INVADER_W_H, y: INVADER_W_H, h: WAVE_HEIGHT * INVADER_W_H, w: WAVE_WIDTH * INVADER_W_H,
	}
	invaders := getInvaders(&waveBoundingBox)

	return &InvaderWave{
		boundingBox:  waveBoundingBox,
		gameBoundary: gameBoundary,
		currentDir:   "LEFT",
		invaders:     invaders,
	}
}

func getInvaders(waveBox *Box) [][]*Invader {
	invaders := [][]*Invader{}
	for i := range WAVE_HEIGHT {
		invaderRow := []*Invader{}
		for j := range WAVE_WIDTH {
			invaderPos := waveBox.getTopLeft().add(Point{x: j * INVADER_W_H, y: i * INVADER_W_H})
			invaderRow = append(invaderRow, NewInvader(invaderPos.x, invaderPos.y))
		}
		invaders = append(invaders, invaderRow)
	}
	return invaders
}

func (w *InvaderWave) BoundingBox() Box {
	return w.boundingBox
}

func (w *InvaderWave) update() {
	w.moveWave()
	// TODO: handle projecting forward later, for now naively move it
}

func (w *InvaderWave) moveWave() {
	yUpdate := 0
	xUpdate := 0

	if w.isAtLateralBoundary() {
		w.currentDir = getOppositeDirection(w.currentDir)
		yUpdate = 1
	}
	xUpdate += getDirScalar(w.currentDir) * X_SPEED

	w.boundingBox.x += xUpdate
	w.boundingBox.y += yUpdate

	w.moveInvaders(xUpdate, yUpdate)
}

func (w *InvaderWave) moveInvaders(x int, y int) {
	for _, invaderRow := range w.invaders {
		for _, invader := range invaderRow {
			invader.moveBy(x, y)
		}
	}
}

func getDirScalar(dir string) int {
	if dir == "LEFT" {
		return -1
	}

	// so gross
	return 1
}

func getOppositeDirection(dir string) string {
	if dir == "LEFT" {
		return "RIGHT"
	}
	return "LEFT"
}

func (w *InvaderWave) isAtLateralBoundary() bool {
	if w.currentDir == "LEFT" && w.boundingBox.x <= w.gameBoundary.leftBorderPos() {
		return true
	}
	if w.currentDir == "RIGHT" && w.boundingBox.x+w.boundingBox.w >= w.gameBoundary.rightBorderPos() {
		return true
	}
	return false
}

func (w *InvaderWave) getUI() []AbstractUiComponent {
	components := []AbstractUiComponent{}
	components = append(components, w.boundingBox.getDebugUI()...)

	for _, invaderRow := range w.invaders {
		for _, invader := range invaderRow {
			components = append(components, invader.getUI()...)
		}
	}

	return components
}
