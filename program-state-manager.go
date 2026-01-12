package main

import (
	"space-invaders/state"
	"space-invaders/ui"
)

type ProgramStateManager struct {
	menu *state.MenuState
	game *GameState
	end  *state.GameOverState
}

func (p *ProgramStateManager) init() {
	p.menu = state.NewMenuState()
}

func (p *ProgramStateManager) update() {
	// TODO: could clean up by making states abstract (Advance(), GetNextState(), GetStaticUI, GetDynamicUI)
	if p.menu != nil {
		res := p.menu.Advance()
		if res == state.EndState() {
			p.game = NewGameState()
			p.menu = nil
		}
		return
	}

	if p.game != nil {
		res := p.game.advance()
		if res == state.EndState() {
			p.end = &state.GameOverState{Score: p.game.scoreTracker.score}
			p.game = nil
		}
		return
	}

	if p.end != nil {
		res := p.end.Advance()
		if res == state.EndState() {
			p.game = NewGameState()
			p.end = nil
		}
		return
	}
}

func (p *ProgramStateManager) GetStaticUI() []ui.StaticUI {
	if p.menu != nil {
		return p.menu.GetStaticUI()
	}

	if p.game != nil {
		return p.game.GetStaticUI()
	}

	if p.end != nil {
		return p.end.GetStaticUI()
	}

	return []ui.StaticUI{}
}

func (p *ProgramStateManager) GetFallbackUI() []ui.StaticUI {
	if p.game == nil {
		return []ui.StaticUI{}
	}

	return p.game.GetFallbackUI()
}

func NewProgramStateMaanger() *ProgramStateManager {
	return &ProgramStateManager{}
}

func (p *ProgramStateManager) GetDynamicUI() []ui.DynamicUI {

	if p.game == nil {
		return []ui.DynamicUI{}
	}

	return p.game.GetDynamicUI()
}
