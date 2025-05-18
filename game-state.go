package main

import (
	"fmt"
)

const X_SPEED = 1

type gameState struct {
	wave         *InvaderWave
	gameBoundary *Box
	player       *Player
}

func (g gameState) advance() AbstractGameState {
	g.wave.update()
	g.player.move() // g.player.update()
	return g        // pointers when?
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
	gameBoundary := Box{x: 0, y: 0, h: 20, w: 100}
	return gameState{
		wave:         NewInvaderWave(&gameBoundary),
		gameBoundary: &gameBoundary,
		player:       NewPlayer(),
	}
}
