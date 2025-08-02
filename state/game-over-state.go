package state

import (
	"fmt"
	"space-invaders/keyboard"
	"space-invaders/ui"
	"space-invaders/utils"
)

type GameOverState struct{ Score int }

func (e *GameOverState) Advance() State {
	if keyboard.GetController().GetCurrentKeypress() == 'r' {
		return EndState()
	}
	return ContinueState()
}

func (e *GameOverState) GetUI() []ui.AbstractUiComponent {
	return []ui.AbstractUiComponent{
		ui.NewSpriteUIComponent(
			fmt.Sprintf("GAME OVER, you scored %d points. Press r to restart or q to quit", e.Score),
			utils.Point{X: 20, Y: 20}),
	}
}
