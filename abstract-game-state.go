package main

type AbstractGameState interface {
	advance() AbstractGameState
	isEnded() bool
	begin()
	getUI() []AbstractUiComponent
}
