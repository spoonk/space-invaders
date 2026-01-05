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
			las.Update()
			g.checkInvaderLaserIntersection()
			if las.Position.Y > g.gameBoundary.Y+g.gameBoundary.H {
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

func (g *GameState) checkInvaderLaserIntersection() {
	for ind, laser := range g.invaderLasers {
		if laser == nil {
			continue
		}

		playerBox := g.player.BoundingBox()

		if playerBox.IsPointWithin(&laser.Position) {
			g.player.RegisterHit()
			g.invaderLasers[ind] = nil
		}
	}
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
	allUI = append(allUI, g.player.GetStaticUI()...)
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
