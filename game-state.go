package main

import (
	"fmt"
	"math/rand"
	"space-invaders/constants"
	"space-invaders/invaders"
	"space-invaders/keyboard"
	"space-invaders/ui"
	"space-invaders/utils"
)

// game state is responsible for managing entities in the game
// the involves removing things, creating things, ending the game, restarting it

type gameState struct {
	wave          *invaders.InvaderWave
	gameBoundary  *utils.Box
	player        *Player
	controller    *keyboard.KeyboardInputController
	activeLaser   *Laser
	scoreTracker  *ScoreTracker
	invaderLasers []*Laser
	debugPane     *DebugPane
}

func (g *gameState) advance() State {
	g.wave.Update()

	if g.wave.IsAtBottom() {
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

		if g.activeLaser.position.Y <= 0 {
			g.activeLaser = nil
		}

		g.checkLaserIntersection()

	}

	if controller.GetCurrentKeypress() == ' ' {
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
			if las.position.Y > g.gameBoundary.Y+g.gameBoundary.H {
				g.invaderLasers[ind] = nil
			}
		}
	}

	// (try to) insert next laser

	if nextLaserInd == -1 {
		return
	}

	for _, row := range g.wave.Invaders {
		for _, inv := range row {
			if !inv.IsDead {
				if rand.Float32() < constants.INVADER_FIRE_PROB {
					// new laser at position
					laserPos := inv.BoundingBox.GetTopLeft().Shifted(1, 1)
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

	if !g.wave.BoundingBox().IsPointWithin(&g.activeLaser.position) {
		return
	}

	// search through invaders, see if laser hits them
	invs := g.wave.Invaders
	for _, row := range invs {
		for _, inv := range row {
			if inv.IsDead {
				continue
			}

			if inv.BoundingBox.IsPointWithin(&g.activeLaser.position) {
				inv.RegisterHit()
				g.scoreTracker.addScore(int(inv.Value))
				g.activeLaser = nil

				if g.wave.NumAliveInvaders() == 0 {
					g.wave = invaders.NewInvaderWave(g.gameBoundary, &utils.Point{X: 0, Y: 0})
				} else {
					g.wave.OnInvaderHit()
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

		if playerBox.IsPointWithin(&laser.position) {
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
	at.X = at.X + 1

	g.activeLaser = NewLaser(&at, -1)
}

func (g *gameState) isEnded() bool {
	return false
}

func (g *gameState) begin() {
	fmt.Println("[game] begin")
}

func (g *gameState) getUI() []ui.AbstractUiComponent {
	var allUI []ui.AbstractUiComponent
	allUI = append(allUI, g.wave.GetUI()...)
	allUI = append(allUI, g.player.getUI()...)
	allUI = append(allUI, ui.GetDebugBoxUI(g.gameBoundary)...)
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
	gameBoundary := utils.Box{
		X: constants.GAME_BOUNDARY.X, Y: constants.GAME_BOUNDARY.Y, H: constants.GAME_BOUNDARY.H, W: constants.GAME_BOUNDARY.W,
	}

	return &gameState{
		wave:         invaders.NewInvaderWave(&gameBoundary, &utils.Point{X: 0, Y: 0}),
		gameBoundary: &gameBoundary,
		player:       NewPlayer(),
		controller:   GetController(),
		scoreTracker: NewScoreTracker(),
		// invaderLasers: []*Laser{nil, nil, nil, nil, nil, nil, nil},
		invaderLasers: make([]*Laser, constants.NUM_INVADER_LASER),
		debugPane:     NewDebugPane(),
	}
}
