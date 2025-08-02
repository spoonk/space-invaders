package main

import (
	"fmt"
	"space-invaders/constants"
	"space-invaders/utils"
)

type DebugPane struct {
	uiPosition utils.Point
	state      *gameState
}

func NewDebugPane() *DebugPane {
	// store sources you want metrics from
	return &DebugPane{
		uiPosition: utils.Point{X: constants.DEBUG_PANE_X, Y: constants.DEBUG_PANE_Y},
	}
}

func (db *DebugPane) getUI(state *gameState) []AbstractUiComponent {
	invaderWavePos := utils.Point{X: state.wave.boundingBox.x, Y: state.wave.boundingBox.y}
	playerPos := state.player.topLeft()
	playerLives := state.player.lives

	components := []AbstractUiComponent{}
	components = append(components, []AbstractUiComponent{
		NewSpriteUIComponent(fmt.Sprintf("wave: (x: %d, y: %d)", invaderWavePos.X, invaderWavePos.Y), db.uiPosition.Shifted(0, 1)),
		NewSpriteUIComponent(fmt.Sprintf("player: (x: %d, y: %d)", playerPos.X, playerPos.Y), db.uiPosition.Shifted(0, 2)),
		NewSpriteUIComponent(fmt.Sprintf("lives remaining: %d", playerLives), db.uiPosition.Shifted(0, 3)),
	}...)

	return components
}
