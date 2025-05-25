package main

import "strconv"

type ScoreTracker struct {
	score int
}

func NewScoreTracker() *ScoreTracker {
	return &ScoreTracker{score: 0}
}

func (s *ScoreTracker) addScore(toAdd int) {
	s.score += toAdd
}

func (s *ScoreTracker) getUI() []AbstractUiComponent {
	return []AbstractUiComponent{
		NewSpriteUIComponent(
			strconv.Itoa(s.score), Point{x: SCORE_X, y: SCORE_Y},
		),
	}
}
