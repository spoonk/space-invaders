package main

// wave designed as a collection of invaders
// wave is responsible for moving all invaders within

type InvaderWave struct {
	boundingBox  Box
	gameBoundary *Box
	invaders     [][]Invader
	currentDir   string
}

const Y_SPEED = 1

// wave needs to define a delta T for how much it should move..
// should be hitching independent: project forward by dt, which can
// lead to you moving > 1 position in a single update

// each wave has the exact same config of invaders
func NewInvaderWave(gameBoundary *Box) *InvaderWave {
	// each wave has the exact same config of invaders
	// todo: instantiate all invaders, compute bounding box
	return &InvaderWave{boundingBox: Box{x: 1, y: 1, h: 4, w: 10}, gameBoundary: gameBoundary, currentDir: "LEFT"}
}

func (w *InvaderWave) BoundingBox() Box {
	return w.boundingBox
}

func (w *InvaderWave) update() {
	w.moveWave()
	// move in current direction
	// TODO: handle projecting forward later, for now naively move it
}

func (w *InvaderWave) moveWave() {

	if w.isAtLateralBoundary() {
		// reverse direction
		w.currentDir = getOppositeDirection(w.currentDir)

		// move down by 1
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
	if w.currentDir == "LEFT" && w.boundingBox.x <= w.gameBoundary.x {
		return true
	}
	if w.currentDir == "RIGHT" && w.boundingBox.x+w.boundingBox.w >= w.gameBoundary.x+w.gameBoundary.w {
		return true
	}
	return false
}

func (w *InvaderWave) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{
		NewSpriteUIComponent("╭", Point{x: w.boundingBox.x, y: w.boundingBox.y}),                                     // top left
		NewSpriteUIComponent("╮", Point{x: w.boundingBox.x + w.boundingBox.w, y: w.boundingBox.y}),                   // top right
		NewSpriteUIComponent("╰", Point{x: w.boundingBox.x, y: w.boundingBox.y + w.boundingBox.h}),                   // bot left
		NewSpriteUIComponent("╯", Point{x: w.boundingBox.x + w.boundingBox.w, y: w.boundingBox.y + w.boundingBox.h}), // bot right
	}
}
