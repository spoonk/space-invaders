package main

// wave designed as a collection of invaders
// wave is responsible for moving all invaders within

type InvaderWave struct {
	boundingBox  Box
	gameBoundary *Box
	invaders     [][]Invader
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
	// each wave has the exact same config of invaders
	// todo: instantiate all invaders, compute bounding box
	return &InvaderWave{
		boundingBox: Box{
			x: INVADER_W_H, y: INVADER_W_H, h: WAVE_HEIGHT * INVADER_W_H, w: WAVE_WIDTH * INVADER_W_H,
		},
		gameBoundary: gameBoundary, currentDir: "LEFT",
	}
}

func (w *InvaderWave) BoundingBox() Box {
	return w.boundingBox
}

func (w *InvaderWave) update() {
	w.moveWave()
	// TODO: handle projecting forward later, for now naively move it
}

func (w *InvaderWave) moveWave() {

	if w.isAtLateralBoundary() {
		w.currentDir = getOppositeDirection(w.currentDir)
		w.boundingBox.y += 1
	}

	w.boundingBox.x += getDirScalar(w.currentDir) * X_SPEED
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
	return w.boundingBox.getDebugUI()
}
