package main

// file containing high-level configuration constants

const DEBUG_BOUNDARY = false

var GAME_BOUNDARY = struct {
	x, y, h, w int
}{
	x: 0, y: 0, h: 30, w: 125,
}

const PLAYER_Y = 30

const (
	INVADER_WAVE_HEIGHT = 5
	INVADER_WAVE_WIDTH  = 11
	X_SPEED             = 1
	Y_SPEED             = 1
	NUM_INVADER_LASER   = 3
)

const (
	NANOSECOND     = 1000000
	FRAME_DURATION = 1000 / 60
)

const (
	DEBUG_PANE_X = 0
	DEBUG_PANE_Y = 31
)

const (
	SCORE_X = 0
	SCORE_Y = 0
)

const (
	GAME_STATE_ID      = "game"
	MENU_STATE_ID      = "menu"
	GAME_OVER_STATE_ID = "game_over"
)

const INVADER_FIRE_PROB = float32(1) / 3000

const PLAYER_W = 3
