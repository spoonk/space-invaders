package main

import (
	"space-invaders/constants"
	"space-invaders/ui"
	"space-invaders/utils"
)

// wave designed as a collection of invaders
// wave is responsible for moving all invaders within

type InvaderWave struct {
	boundingBox  Box
	gameBoundary *Box
	invaders     [][]*Invader
	currentDir   string
}

// TODO: different types of invaders
// TODO: resize bounding box when invader dies

// wave needs to define a delta T for how much it should move..
// should be hitching independent: project forward by dt, which can
// lead to you moving > 1 position in a single update

// each wave has the exact same config of invaders
func NewInvaderWave(gameBoundary *Box, startPoint *utils.Point) *InvaderWave {
	if startPoint == nil {
		startPoint = &utils.Point{X: 0, Y: 0}
	}
	invaders := getInvaders(startPoint)

	waveBoundingBox := inferBoundingBox(gameBoundary, invaders)

	return &InvaderWave{
		boundingBox:  waveBoundingBox,
		gameBoundary: gameBoundary,
		currentDir:   "LEFT",
		invaders:     invaders,
	}
}

func getInvaders(topLeft *utils.Point) [][]*Invader {
	invaders := [][]*Invader{}
	for i := range constants.INVADER_WAVE_HEIGHT {
		invaderRow := []*Invader{}
		for j := range constants.INVADER_WAVE_WIDTH {
			invaderPos := topLeft.Add(utils.Point{X: j * (INVADER_W + 2), Y: i * (INVADER_H + 1)})
			invaderRow = append(invaderRow, NewInvader(invaderPos.X, invaderPos.Y))
		}
		invaders = append(invaders, invaderRow)
	}
	return invaders
}

func inferBoundingBox(gameBoundary *Box, invaders [][]*Invader) Box {
	// assertion: at least one invader alive
	minX := gameBoundary.x + gameBoundary.w
	minY := gameBoundary.y + gameBoundary.h
	maxX := 0
	maxY := 0

	for _, invaderRow := range invaders {
		for _, invader := range invaderRow {
			if invader.isDead {
				continue
			}

			minX = min(minX, invader.topLeft().X)
			maxX = max(maxX, invader.topLeft().X+INVADER_W)
			minY = min(minY, invader.topLeft().Y)
			maxY = max(maxY, invader.topLeft().Y+INVADER_H)
		}
	}

	return Box{x: minX, y: minY, w: maxX - minX, h: maxY - minY}
}

func (w *InvaderWave) update() {
	w.moveWave()
	// TODO: handle projecting forward later, for now naively move it
}

func (w *InvaderWave) moveWave() {
	// todo: if I want to cache the areAllInvadersDead computation, would set dirty bit here
	yUpdate := 0
	xUpdate := 0

	if w.isAtLateralBoundary() {
		w.currentDir = getOppositeDirection(w.currentDir)
		yUpdate = 1
	}
	xUpdate += getDirScalar(w.currentDir) * constants.X_SPEED

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

func (w *InvaderWave) onInvaderHit() {
	// update boundingbox
	w.boundingBox = inferBoundingBox(w.gameBoundary, w.invaders)
}

func (w *InvaderWave) numAliveInvaders() int {
	cnt := 0
	for _, row := range w.invaders {
		for _, inv := range row {
			if !inv.isDead {
				cnt++
			}
		}
	}
	return cnt
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

// TODO: these should be box functions
func (w *InvaderWave) isAtLateralBoundary() bool {
	if w.currentDir == "LEFT" && w.boundingBox.x <= w.gameBoundary.leftBorderPos() {
		return true
	}
	if w.currentDir == "RIGHT" && w.boundingBox.x+w.boundingBox.w >= w.gameBoundary.rightBorderPos() {
		return true
	}
	return false
}

func (w *InvaderWave) isAtBottom() bool {
	return w.boundingBox.y+w.boundingBox.h >= w.gameBoundary.y+w.gameBoundary.h
}

func (w *InvaderWave) BoundingBox() Box {
	return w.boundingBox
}

func (w *InvaderWave) getUI() []ui.AbstractUiComponent {
	components := []ui.AbstractUiComponent{}
	components = append(components, w.boundingBox.getDebugUI()...)

	for _, invaderRow := range w.invaders {
		for _, invader := range invaderRow {
			components = append(components, invader.getUI()...)
		}
	}

	return components
}
