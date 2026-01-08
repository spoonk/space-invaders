package state

import (
	"fmt"
	"space-invaders/keyboard"
	"space-invaders/ui"
)

type GameOverState struct{ Score int }

func (e *GameOverState) Advance() State {
	if keyboard.GetController().GetCurrentKeypress() == 'r' {
		return EndState()
	}
	return ContinueState()
}

func (e *GameOverState) GetStaticUI() []ui.StaticUI {
	return []ui.StaticUI{ui.NewCenteredTextUIComponent(fmt.Sprintf("GAME OVER, you scored %d points. Press r to restart or q to quit", e.Score))}
}
