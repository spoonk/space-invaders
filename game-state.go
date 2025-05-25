package main

import (
	"fmt"
)

// game state is responsible for managing entities in the game
// the involves removing things, creating things, ending the game, restarting it

type gameState struct {
	wave         *InvaderWave
	gameBoundary *Box
	player       *Player
	controller   *KeyboardInputController
	activeLaser  *Laser
	scoreTracker *ScoreTracker
}

func (g *gameState) advance() State {
	g.wave.update()

	if g.wave.isAtBottom() {
		// TODO: transition to game over screen
		return EndState()
	}
	g.player.move() // g.player.update()

	if g.activeLaser != nil {
		g.activeLaser.update()

		if g.activeLaser.position.y <= 0 {
			g.activeLaser = nil
		}

		g.checkLaserIntersection()
	}

	if controller.getCurrentKeypress() == ' ' {
		g.handleShoot()
	}

	return ContinueState()
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
				g.scoreTracker.addScore(int(inv.value))
				g.activeLaser = nil

				if g.wave.areAllInvadersDead() {
					g.wave = NewInvaderWave(g.gameBoundary, &Point{x: 0, y: 0})
				} else {
					g.wave.onInvaderHit()
				}

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
	allUI = append(allUI, g.gameBoundary.getDebugUI()...)
	allUI = append(allUI, g.scoreTracker.getUI()...)

	if g.activeLaser != nil {
		allUI = append(allUI, g.activeLaser.getUI()...)
	}

	return allUI
}

func NewGameState() *gameState {
	gameBoundary := Box{
		x: GAME_BOUNDARY.x, y: GAME_BOUNDARY.y, h: GAME_BOUNDARY.h, w: GAME_BOUNDARY.w,
	}

	return &gameState{
		wave:         NewInvaderWave(&gameBoundary, &Point{x: 0, y: 0}),
		gameBoundary: &gameBoundary,
		player:       NewPlayer(),
		controller:   GetController(),
		scoreTracker: NewScoreTracker(),
	}
}
