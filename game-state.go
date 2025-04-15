package main

import (
	"fmt"
	"math/rand/v2"
)

const X_SPEED = 1

// hmmm right now everything needs to know where it is
// which is kinda sus?

type gameState struct {
	wave          Wave
	gameBoundary  Box
	waveDirection int
	player        *Player
}

func (g gameState) advance() AbstractGameState {
	dir, bumped := g.nextDirection()
	g.waveDirection = dir
	var nextWave = getNextWavePos(g.wave.boundingBox, g.waveDirection, bumped)
	g.wave.boundingBox = nextWave
	g.movePlayer()
	return g // pointers when?
}

func (g gameState) movePlayer() {
	g.player.moveTo(Point{x: rand.IntN(100), y: rand.IntN(100)})
}

func getNextWavePos(currWavePos Box, direction int, bumped bool) Box {
	yShift := 0
	if bumped {
		yShift = 1
	}

	xShift := direction * X_SPEED
	if bumped {
		xShift = 0
	}

	return Box{x: currWavePos.x + xShift, y: currWavePos.y + yShift, h: currWavePos.h, w: currWavePos.w}
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
	var allUI []AbstractUiComponent
	allUI = append(allUI, g.wave.getUI()...)
	allUI = append(allUI, g.player.getUI()...)
	return allUI
}

func NewGameState() AbstractGameState {
	return gameState{
		wave:          Wave{boundingBox: Box{x: 1, y: 1, h: 4, w: 10}},
		gameBoundary:  Box{x: 0, y: 0, h: 20, w: 100},
		waveDirection: 1,
		player:        NewPlayer(),
	}
}
