package main

import (
	"fmt"
	"math/rand"
	"space-invaders/constants"
)

// game state is responsible for managing entities in the game
// the involves removing things, creating things, ending the game, restarting it

type gameState struct {
	wave          *InvaderWave
	gameBoundary  *Box
	player        *Player
	controller    *KeyboardInputController
	activeLaser   *Laser
	scoreTracker  *ScoreTracker
	invaderLasers []*Laser
	debugPane     *DebugPane
}

func (g *gameState) advance() State {
	g.wave.update()

	if g.wave.isAtBottom() {
		return EndState()
	}
	g.updateLaser()
	g.updateInvaderLasers()

	if g.player.lives == 0 {
		return EndState()
	}
	g.player.move() // g.player.update()

	return ContinueState()
}

// figure out how to shoot invaders

func (g *gameState) updateLaser() {
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
}

func (g *gameState) updateInvaderLasers() {
	// we have 55 invaders, that's 55 entities updating per loop
	// the probability we fire should be proportional to the number of non-dead invaders
	// then just pick some scalar and tune it lol

	// can achieve this probability by looping through invaders & deciding to fire for each one

	// need to update current lasers and determine whether or not to fire a new one
	nextLaserInd := -1

	for ind, las := range g.invaderLasers {
		if las == nil {
			nextLaserInd = ind
		} else {
			las.update()
			g.checkInvaderLaserIntersection()
			if las.position.y > g.gameBoundary.y+g.gameBoundary.h {
				g.invaderLasers[ind] = nil
			}
		}
	}

	// (try to) insert next laser

	if nextLaserInd == -1 {
		return
	}

	for _, row := range g.wave.invaders {
		for _, inv := range row {
			if !inv.isDead {
				if rand.Float32() < constants.INVADER_FIRE_PROB {
					// new laser at position
					laserPos := inv.boundingBox.getTopLeft().shifted(1, 1)
					g.invaderLasers[nextLaserInd] = NewLaser(&laserPos, 1)
				}
			}
		}
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
				g.scoreTracker.addScore(int(inv.value))
				g.activeLaser = nil

				if g.wave.numAliveInvaders() == 0 {
					g.wave = NewInvaderWave(g.gameBoundary, &Point{x: 0, y: 0})
				} else {
					g.wave.onInvaderHit()
				}

				return
			}
		}
	}
}

func (g *gameState) checkInvaderLaserIntersection() {
	for ind, laser := range g.invaderLasers {
		if laser == nil {
			continue
		}

		playerBox := g.player.boundingBox()

		if playerBox.isPointWithin(&laser.position) {
			g.player.registerHit()
			g.invaderLasers[ind] = nil
		}
	}
}

func (g *gameState) handleShoot() {
	if g.activeLaser != nil {
		return
	}

	at := g.player.pos
	at.x = at.x + 1

	g.activeLaser = NewLaser(&at, -1)
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

	for _, invLaser := range g.invaderLasers {
		if invLaser != nil {
			allUI = append(allUI, invLaser.getUI()...)
		}
	}

	allUI = append(allUI, g.debugPane.getUI(g)...)

	return allUI
}

func NewGameState() *gameState {
	gameBoundary := Box{
		x: constants.GAME_BOUNDARY.X, y: constants.GAME_BOUNDARY.Y, h: constants.GAME_BOUNDARY.H, w: constants.GAME_BOUNDARY.W,
	}

	return &gameState{
		wave:         NewInvaderWave(&gameBoundary, &Point{x: 0, y: 0}),
		gameBoundary: &gameBoundary,
		player:       NewPlayer(),
		controller:   GetController(),
		scoreTracker: NewScoreTracker(),
		// invaderLasers: []*Laser{nil, nil, nil, nil, nil, nil, nil},
		invaderLasers: make([]*Laser, constants.NUM_INVADER_LASER),
		debugPane:     NewDebugPane(),
	}
}
