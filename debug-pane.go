package main

import (
	"fmt"
	"space-invaders/constants"
	"space-invaders/ui"
	"space-invaders/utils"
)

type DebugPane struct {
	uiPosition utils.Point
	state      *GameState
}

func NewDebugPane() *DebugPane {
	// store sources you want metrics from
	return &DebugPane{
		uiPosition: utils.Point{X: constants.DEBUG_PANE_X, Y: constants.DEBUG_PANE_Y},
	}
}

func (db *DebugPane) GetStaticUI(state *GameState) []ui.StaticUI {
	invaderWavePos := utils.Point{X: state.wave.BoundingBox().X, Y: state.wave.BoundingBox().Y}
	playerPos := state.player.TopLeft()
	playerLives := state.player.Lives

	components := []ui.StaticUI{}
	components = append(components, []ui.StaticUI{
		ui.NewSpriteUIComponent(fmt.Sprintf("wave: (x: %d, y: %d)", invaderWavePos.X, invaderWavePos.Y), db.uiPosition.Shifted(0, 1)),
		ui.NewSpriteUIComponent(fmt.Sprintf("player: (x: %d, y: %d)", playerPos.X, playerPos.Y), db.uiPosition.Shifted(0, 2)),
		ui.NewSpriteUIComponent(fmt.Sprintf("lives remaining: %d", playerLives), db.uiPosition.Shifted(0, 3)),
	}...)

	return components
}
