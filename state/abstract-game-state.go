package state

import "space-invaders/ui"

var STATE_TRANSITIONS = struct {
	END, CONTINUE int
}{
	END:      0,
	CONTINUE: 1,
}

type (
	State int
)

func EndState() State {
	return State(STATE_TRANSITIONS.END)
}

func ContinueState() State {
	return State(STATE_TRANSITIONS.CONTINUE)
}

type AbstractGameState interface {
	advance() State
	isEnded() bool
	begin()
	GetUI() []ui.AbstractUiComponent
}
