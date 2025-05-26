package main

import "fmt"

type GameOverState struct{ score int }

func (e *GameOverState) advance() State {
	if GetController().getCurrentKeypress() == 'r' {
		return EndState()
	}
	return ContinueState()
}

func (e *GameOverState) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{
		NewSpriteUIComponent(
			fmt.Sprintf("GAME OVER, you scored %d points. Press r to restart or q to quit", e.score),
			Point{x: 20, y: 20}),
	}
}
