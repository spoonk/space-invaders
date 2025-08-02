package main

import "space-invaders/utils"

type MenuState struct{}

func (m *MenuState) advance() State {
	if GetController().getCurrentKeypress() == '1' {
		return EndState()
	}

	return ContinueState()
}

func NewMenuState() *MenuState {
	return &MenuState{}
}

func (m *MenuState) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{
		NewSpriteUIComponent(
			"Welcome to space invaders! Press 1 to play",
			utils.Point{X: 20, Y: 20}),
	}
}
