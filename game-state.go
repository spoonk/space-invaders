package main

import (
	"fmt"
)

const X_SPEED = 1

type box struct {
	x, y, h, w int
}

func (b box) leftBorderPos() int {
	return b.x
}

func (b box) rightBorderPos() int {
	return b.x + b.w
}

type gameState struct {
	wave          box
	gameBoundary  box
	waveDirection int
}

func (g gameState) advance() AbstractGameState {
	dir, bumped := g.nextDirection()
	g.waveDirection = dir
	var nextWave = getNextWavePos(g.wave, g.waveDirection, bumped)
	g.wave = nextWave

	return g
}

func getNextWavePos(wave box, direction int, bumped bool) box {
	yShift := 0
	if bumped {
		yShift = 1
	}

	xShift := direction * X_SPEED
	if bumped {
		xShift = 0
	}

	return box{x: wave.x + xShift, y: wave.y + yShift, h: wave.h, w: wave.w}
}

func (g gameState) nextDirection() (int, bool) {
	if g.waveDirection == 1 {
		if g.wave.rightBorderPos() >= g.gameBoundary.rightBorderPos() {
			return -1, true
		}
	} else {

		if g.wave.leftBorderPos() <= g.gameBoundary.leftBorderPos() {
			return 1, true
		}
	}
	return g.waveDirection, false
}

func (g gameState) isEnded() bool {
	return false
}

func (g gameState) begin() {
	fmt.Println("[game] begin")
}

func (g gameState) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{
		NewSpriteUIComponent("╭", Point{x: g.wave.x, y: g.wave.y}),                       // top left
		NewSpriteUIComponent("╮", Point{x: g.wave.x + g.wave.w, y: g.wave.y}),            // top right
		NewSpriteUIComponent("╰", Point{x: g.wave.x, y: g.wave.y + g.wave.h}),            // bot left
		NewSpriteUIComponent("╯", Point{x: g.wave.x + g.wave.w, y: g.wave.y + g.wave.h}), // bot right
	}
}

func NewGameState() AbstractGameState {
	return gameState{
		wave:          box{x: 1, y: 1, h: 4, w: 10},
		gameBoundary:  box{x: 0, y: 0, h: 20, w: 100},
		waveDirection: 1,
	}
}
