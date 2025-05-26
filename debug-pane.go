package main

import (
	"fmt"
)

type DebugPane struct {
	uiPosition Point
	state      *gameState
}

func NewDebugPane() *DebugPane {
	// store sources you want metrics from
	return &DebugPane{
		uiPosition: Point{x: DEBUG_PANE_X, y: DEBUG_PANE_Y},
	}
}

func (db *DebugPane) getUI(state *gameState) []AbstractUiComponent {
	invaderWavePos := Point{x: state.wave.boundingBox.x, y: state.wave.boundingBox.y}
	playerPos := state.player.topLeft()
	playerLives := state.player.lives

	components := []AbstractUiComponent{}
	components = append(components, []AbstractUiComponent{
		NewSpriteUIComponent(fmt.Sprintf("wave: (x: %d, y: %d)", invaderWavePos.x, invaderWavePos.y), db.uiPosition.shifted(0, 1)),
		NewSpriteUIComponent(fmt.Sprintf("player: (x: %d, y: %d)", playerPos.x, playerPos.y), db.uiPosition.shifted(0, 2)),
		NewSpriteUIComponent(fmt.Sprintf("lives remaining: %d", playerLives), db.uiPosition.shifted(0, 3)),
	}...)

	return components
}
