package main

type AbstractGameState interface {
	advance()
	isEnded() bool
	begin()
	getUI() []AbstractUiComponent
}
