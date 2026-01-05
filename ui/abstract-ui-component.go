package ui

import (
	"space-invaders/utils"
)

type AbstractUiComponent interface {
	GetTopLeft() utils.Point
	GetRasterized() []string
}

type StaticUI interface {
	GetTopLeft() utils.Point
	GetRasterized() []string
}

type DynamicUI interface {
	GetBoundingBox() utils.Box
	GetPath() string
}

type HydratedDynamicUI struct {
	BoundingBox utils.Box
	Image       *[][]float64
}
