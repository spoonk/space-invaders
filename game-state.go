package main

import (
	"fmt"
)

const X_SPEED = 1
const DEBUG_BOUNDARY = true

// game state is responsible for managing entities in the game
// the involves removing things, creating things, ending the game, restarting it

type gameState struct {
	wave         *InvaderWave
	gameBoundary *Box
	player       *Player
}

func (g *gameState) advance() {
	g.wave.update()
	g.player.move() // g.player.update()
}

func (g *gameState) isEnded() bool {
	return false
}

func (g *gameState) begin() {
	fmt.Println("[game] begin")
}

// TODO: render the game boundary as a box
func (g *gameState) getUI() []AbstractUiComponent {
	var allUI []AbstractUiComponent
	allUI = append(allUI, g.wave.getUI()...)
	allUI = append(allUI, g.player.getUI()...)
	if DEBUG_BOUNDARY {
		allUI = append(allUI, g.gameBoundary.getDebugUI()...)
	}
	return allUI
}

func NewGameState() AbstractGameState {
	gameBoundary := Box{x: 0, y: 0, h: 100, w: 125}
	return &gameState{
		wave:         NewInvaderWave(&gameBoundary),
		gameBoundary: &gameBoundary,
		player:       NewPlayer(),
	}
}
