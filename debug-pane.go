package main

import (
	"fmt"
)

type DebugPane struct {
	uiPosition Point
	state      *gameState
}

func NewDebugPane(g *gameState) *DebugPane {
	// store sources you want metrics from
	return &DebugPane{
		uiPosition: Point{x: DEBUG_PANE_X, y: DEBUG_PANE_Y},

		state: g,
	}
}

func (db *DebugPane) getUI() []AbstractUiComponent {
	invaderWavePos := Point{x: db.state.wave.boundingBox.x, y: db.state.wave.boundingBox.y}
	playerPos := db.state.player.topLeft()

	components := []AbstractUiComponent{}
	components = append(components, []AbstractUiComponent{
		NewSpriteUIComponent(fmt.Sprintf("wave: (x: %d, y: %d)", invaderWavePos.x, invaderWavePos.y), db.uiPosition.shifted(0, 1)),
		NewSpriteUIComponent(fmt.Sprintf("player: (x: %d, y: %d)", playerPos.x, playerPos.y), db.uiPosition.shifted(0, 2)),
	}...)

	return components
}
