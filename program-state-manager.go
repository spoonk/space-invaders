package main

type ProgramStateManager struct {
	game *gameState
}

func (p *ProgramStateManager) init() {
	p.game = NewGameState()
}

func (p *ProgramStateManager) update() {
	if p.game == nil {
		return
	}

	res := p.game.advance()
	if res == EndState() {
		p.game = nil
		// Run = false
	}
}

func (p *ProgramStateManager) getUI() []AbstractUiComponent {
	if p.game == nil {
		return []AbstractUiComponent{}
	}
	return p.game.getUI()
}

func NewProgramStateMaanger() *ProgramStateManager {
	return &ProgramStateManager{}
}

// define
