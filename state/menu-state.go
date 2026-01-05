package state

import (
	"space-invaders/keyboard"
	"space-invaders/ui"
	"space-invaders/utils"
)

type MenuState struct{}

func (m *MenuState) Advance() State {
	if keyboard.GetController().GetCurrentKeypress() == '1' {
		return EndState()
	}

	return ContinueState()
}

func NewMenuState() *MenuState {
	return &MenuState{}
}

func (m *MenuState) GetStaticUI() []ui.StaticUI {
	return []ui.StaticUI{
		ui.NewSpriteUIComponent(
			"Welcome to space invaders! Press 1 to play",
			utils.Point{X: 20, Y: 20}),
	}
}
