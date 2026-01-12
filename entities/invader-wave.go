package entities

import (
	"space-invaders/constants"
	"space-invaders/ui"
	"space-invaders/utils"
)

// wave designed as a collection of invaders
// wave is responsible for moving all invaders within

type InvaderWave struct {
	boundingBox  utils.Box
	gameBoundary *utils.Box
	Invaders     [][]*Invader
	currentDir   string
}

// TODO: different types of invaders
// TODO: resize bounding box when invader dies

// wave needs to define a delta T for how much it should move..
// should be hitching independent: project forward by dt, which can
// lead to you moving > 1 position in a single update

// each wave has the exact same config of invaders
func NewInvaderWave(gameBoundary *utils.Box, startPoint *utils.Point) *InvaderWave {
	if startPoint == nil {
		startPoint = &utils.Point{X: 0, Y: 0}
	}
	invaders := getInvaders(startPoint)

	waveBoundingBox := inferBoundingBox(gameBoundary, invaders)

	return &InvaderWave{
		boundingBox:  waveBoundingBox,
		gameBoundary: gameBoundary,
		currentDir:   "LEFT",
		Invaders:     invaders,
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

func inferBoundingBox(gameBoundary *utils.Box, invaders [][]*Invader) utils.Box {
	// assertion: at least one invader alive
	minX := gameBoundary.X + gameBoundary.W
	minY := gameBoundary.Y + gameBoundary.H
	maxX := 0
	maxY := 0

	for _, invaderRow := range invaders {
		for _, invader := range invaderRow {
			if invader.IsDead {
				continue
			}

			minX = min(minX, invader.topLeft().X)
			maxX = max(maxX, invader.topLeft().X+INVADER_W)
			minY = min(minY, invader.topLeft().Y)
			maxY = max(maxY, invader.topLeft().Y+INVADER_H)
		}
	}

	return utils.Box{X: minX, Y: minY, W: maxX - minX, H: maxY - minY}
}

func (w *InvaderWave) Update() {
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

	w.boundingBox.X += xUpdate
	w.boundingBox.Y += yUpdate

	w.MoveInvaders(xUpdate, yUpdate)
}

func (w *InvaderWave) MoveInvaders(x int, y int) {
	for _, invaderRow := range w.Invaders {
		for _, invader := range invaderRow {
			invader.moveBy(x, y)
		}
	}
}

func (w *InvaderWave) OnInvaderHit() {
	// update boundingbox
	w.boundingBox = inferBoundingBox(w.gameBoundary, w.Invaders)
}

func (w *InvaderWave) NumAliveInvaders() int {
	cnt := 0
	for _, row := range w.Invaders {
		for _, inv := range row {
			if !inv.IsDead {
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
	if w.currentDir == "LEFT" && w.boundingBox.X <= w.gameBoundary.LeftBorderPos() {
		return true
	}
	if w.currentDir == "RIGHT" && w.boundingBox.X+w.boundingBox.W >= w.gameBoundary.RightBorderPos() {
		return true
	}
	return false
}

func (w *InvaderWave) IsAtBottom() bool {
	return w.boundingBox.Y+w.boundingBox.H >= w.gameBoundary.Y+w.gameBoundary.H
}

func (w *InvaderWave) BoundingBox() utils.Box {
	return w.boundingBox
}

func (w *InvaderWave) GetStaticUI() []ui.StaticUI {
	components := []ui.StaticUI{}
	components = append(components, ui.GetDebugBoxUI(&w.boundingBox)...)

	return components
}

func (w *InvaderWave) GetDynamicUI() []ui.DynamicUI {
	components := []ui.DynamicUI{}

	for _, invaderRow := range w.Invaders {
		for _, invader := range invaderRow {
			components = append(components, invader.GetDynamicUI()...)
		}
	}

	return components
}

func (w *InvaderWave) GetFallbackUI() []ui.StaticUI {
	components := []ui.StaticUI{}

	for _, invaderRow := range w.Invaders {
		for _, invader := range invaderRow {
			components = append(components, invader.GetFallbackUI()...)
		}
	}

	return components
}
