package main

import (
	"space-invaders/constants"
	"space-invaders/ui"
	"space-invaders/utils"
	"strconv"
)

type ScoreTracker struct {
	score int
}

func NewScoreTracker() *ScoreTracker {
	return &ScoreTracker{score: 0}
}

func (s *ScoreTracker) addScore(toAdd int) {
	s.score += toAdd
}

func (s *ScoreTracker) GetUI() []ui.AbstractUiComponent {
	return []ui.AbstractUiComponent{
		ui.NewSpriteUIComponent(
			strconv.Itoa(s.score), utils.Point{X: constants.SCORE_X, Y: constants.SCORE_Y},
		),
	}
}
