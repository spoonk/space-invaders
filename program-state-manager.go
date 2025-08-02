package main

import "space-invaders/ui"

type ProgramStateManager struct {
	menu *MenuState
	game *gameState
	end  *GameOverState
}

func (p *ProgramStateManager) init() {
	p.menu = NewMenuState()
}

func (p *ProgramStateManager) update() {
	// right now I have to special case everything....
	if p.menu != nil {
		res := p.menu.advance()
		if res == EndState() {
			p.game = NewGameState()
			p.menu = nil
		}
		return
	}

	if p.game != nil {
		res := p.game.advance()
		if res == EndState() {
			p.end = &GameOverState{score: p.game.scoreTracker.score}
			p.game = nil
		}
		return
	}

	if p.end != nil {
		res := p.end.advance()
		if res == EndState() {
			p.game = NewGameState()
			p.end = nil
		}
		return
	}
}

func (p *ProgramStateManager) getUI() []ui.AbstractUiComponent {
	if p.menu != nil {
		return p.menu.getUI()
	}

	if p.game != nil {
		return p.game.getUI()
	}

	if p.end != nil {
		return p.end.getUI()
	}

	return []ui.AbstractUiComponent{}
	// return p.game.getUI()
}

func NewProgramStateMaanger() *ProgramStateManager {
	return &ProgramStateManager{}
}

// define
