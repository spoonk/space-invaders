package main

import (
	"fmt"
	"math/rand"
	"space-invaders/constants"
	"space-invaders/entities"
	"space-invaders/keyboard"
	"space-invaders/state"
	"space-invaders/ui"
	"space-invaders/utils"
)

// game state is responsible for managing entities in the game
// the involves removing things, creating things, ending the game, restarting it

type GameState struct {
	wave          *entities.InvaderWave
	gameBoundary  *utils.Box
	player        *entities.Player
	controller    *keyboard.KeyboardInputController
	activeLaser   *entities.Laser
	scoreTracker  *ScoreTracker
	invaderLasers []*entities.Laser
	debugPane     *DebugPane
}

func (g *GameState) advance() state.State {
	g.wave.Update()

	if g.wave.IsAtBottom() {
		return state.EndState()
	}

	g.updatePlayerLaser()
	g.updateInvaderLasers()

	if g.player.Lives == 0 {
		return state.EndState()
	}

	g.player.Move()

	return state.ContinueState()
}

// figure out how to shoot invaders

func (g *GameState) updatePlayerLaser() {
	if g.activeLaser != nil {
		g.activeLaser.Update()

		if g.activeLaser.Position.Y <= 0 {
			g.activeLaser = nil
		}

		g.checkLaserIntersection()
	}

	if g.controller.GetCurrentKeypress() == ' ' {
		g.handleShoot()
	}
}

func (g *GameState) updateInvaderLasers() {
	nextLaserInd := -1

	for ind, currLaser := range g.invaderLasers {
		if currLaser == nil {
			nextLaserInd = ind
		} else {
			currLaser.Update()
			if g.checkInvaderLaserIntersection(currLaser) {
				g.player.RegisterHit()
				g.invaderLasers[ind] = nil
			}
			if currLaser.Position.Y > g.gameBoundary.Y+g.gameBoundary.H {
				g.invaderLasers[ind] = nil
			}
		}
	}

	// all laser slots currently in use
	if nextLaserInd == -1 {
		return
	}

	// Determine who shoots next
	for _, row := range g.wave.Invaders {
		for _, inv := range row {
			if !inv.IsDead {
				// note: this is pretty biased towards invaders in the top-left
				// it's very unlikely that an invader in the bottom right will ever have a chance to fire a laser
				// TODO: shuffle invaders, then evaluate probabilities
				if rand.Float32() < constants.INVADER_FIRE_PROB {
					laserPos := inv.BoundingBox.GetTopLeft().Shifted(1, 1)
					g.invaderLasers[nextLaserInd] = entities.NewLaser(&laserPos, 1)
				}
			}
		}
	}
}

func (g *GameState) checkLaserIntersection() {
	if g.activeLaser == nil {
		return
	}

	if !g.wave.BoundingBox().IsPointWithin(&g.activeLaser.Position) {
		return
	}

	// search through invaders, see if laser hits them
	invs := g.wave.Invaders
	for _, row := range invs {
		for _, inv := range row {
			if inv.IsDead {
				continue
			}

			if inv.BoundingBox.IsPointWithin(&g.activeLaser.Position) {
				inv.RegisterHit()
				g.scoreTracker.addScore(int(inv.Value))
				g.activeLaser = nil

				if g.wave.NumAliveInvaders() == 0 {
					g.wave = entities.NewInvaderWave(g.gameBoundary, &utils.Point{X: 0, Y: 0})
				} else {
					g.wave.OnInvaderHit()
				}

				return
			}
		}
	}
}

func (g *GameState) checkInvaderLaserIntersection(laser *entities.Laser) bool {
	playerBox := g.player.BoundingBox()
	return playerBox.IsPointWithin(&laser.Position)
}

func (g *GameState) handleShoot() {
	if g.activeLaser != nil {
		return
	}

	at := g.player.Pos
	at.X = at.X + 1

	g.activeLaser = entities.NewLaser(&at, -1)
}

func (g *GameState) isEnded() bool {
	return false
}

func (g *GameState) begin() {
	fmt.Println("[game] begin")
}

func (g *GameState) GetStaticUI() []ui.StaticUI {
	var allUI []ui.StaticUI

	allUI = append(allUI, g.wave.GetStaticUI()...)
	// allUI = append(allUI, g.player.GetStaticUI()...)
	allUI = append(allUI, ui.GetDebugBoxUI(g.gameBoundary)...)
	allUI = append(allUI, g.scoreTracker.GetStaticUI()...)

	if g.activeLaser != nil {
		allUI = append(allUI, g.activeLaser.GetStaticUI()...)
	}

	for _, invLaser := range g.invaderLasers {
		if invLaser != nil {
			allUI = append(allUI, invLaser.GetStaticUI()...)
		}
	}

	allUI = append(allUI, g.debugPane.GetStaticUI(g)...)

	return allUI
}

func (g *GameState) GetDynamicUI() []ui.DynamicUI {
	allUI := []ui.DynamicUI{}

	allUI = append(allUI, g.wave.GetDynamicUI()...)
	allUI = append(allUI, g.player.GetDynamicUI()...)

	return allUI

}
func NewGameState() *GameState {
	gameBoundary := utils.Box{
		X: constants.GAME_BOUNDARY.X, Y: constants.GAME_BOUNDARY.Y, H: constants.GAME_BOUNDARY.H, W: constants.GAME_BOUNDARY.W,
	}

	return &GameState{
		wave:          entities.NewInvaderWave(&gameBoundary, &utils.Point{X: 0, Y: 0}),
		gameBoundary:  &gameBoundary,
		player:        entities.NewPlayer(),
		controller:    keyboard.GetController(),
		scoreTracker:  NewScoreTracker(),
		invaderLasers: make([]*entities.Laser, constants.NUM_INVADER_LASER),
		debugPane:     NewDebugPane(),
	}
}
