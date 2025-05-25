package main

import (
	"fmt"
)

const X_SPEED = 1
const DEBUG_BOUNDARY = false

// game state is responsible for managing entities in the game
// the involves removing things, creating things, ending the game, restarting it

type gameState struct {
	wave         *InvaderWave
	gameBoundary *Box
	player       *Player
	controller   *KeyboardInputController
	activeLaser  *Laser
}

func (g *gameState) advance() {
	g.wave.update()
	g.player.move() // g.player.update()

	// todo: check if laser will hit invader

	if g.activeLaser != nil {
		g.activeLaser.update()

		if g.activeLaser.position.y <= 0 {
			g.activeLaser = nil
		}

		// check if laser now on top of invader
		g.checkLaserIntersection()
	}

	// if space pressed and no laser, spawn new laser
	if controller.getCurrentKeypress() == ' ' {
		g.handleShoot()
	}

}

func (g *gameState) checkLaserIntersection() {
	if g.activeLaser == nil {
		return
	}

	if !g.wave.boundingBox.isPointWithin(&g.activeLaser.position) {
		return
	}

	// search through invaders, see if laser hits them
	invaders := g.wave.invaders
	for _, row := range invaders {
		for _, inv := range row {
			if inv.isDead {
				continue
			}

			if inv.boundingBox.isPointWithin(&g.activeLaser.position) {
				inv.registerHit()
				g.activeLaser = nil
				return
			}
		}
	}
}

func (g *gameState) handleShoot() {
	if g.activeLaser != nil {
		return
	}

	at := g.player.pos
	at.x = at.x + 1

	g.activeLaser = NewLaser(&at)
}

func (g *gameState) isEnded() bool {
	return false
}

func (g *gameState) begin() {
	fmt.Println("[game] begin")
}

func (g *gameState) getUI() []AbstractUiComponent {
	var allUI []AbstractUiComponent
	allUI = append(allUI, g.wave.getUI()...)
	allUI = append(allUI, g.player.getUI()...)
	if DEBUG_BOUNDARY {
		allUI = append(allUI, g.gameBoundary.getDebugUI()...)
	}

	if g.activeLaser != nil {
		allUI = append(allUI, g.activeLaser.getUI()...)
	}
	return allUI
}

func NewGameState() *gameState {
	gameBoundary := Box{x: 0, y: 0, h: 30, w: 125}
	return &gameState{
		wave:         NewInvaderWave(&gameBoundary, &Point{x: 0, y: 0}),
		gameBoundary: &gameBoundary,
		player:       NewPlayer(),
		controller:   GetController(),
	}
}
