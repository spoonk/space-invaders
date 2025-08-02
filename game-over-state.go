package main

import (
	"fmt"
	"space-invaders/ui"
	"space-invaders/utils"
)

type GameOverState struct{ score int }

func (e *GameOverState) advance() State {
	if GetController().getCurrentKeypress() == 'r' {
		return EndState()
	}
	return ContinueState()
}

func (e *GameOverState) getUI() []ui.AbstractUiComponent {
	return []ui.AbstractUiComponent{
		ui.NewSpriteUIComponent(
			fmt.Sprintf("GAME OVER, you scored %d points. Press r to restart or q to quit", e.score),
			utils.Point{X: 20, Y: 20}),
	}
}
