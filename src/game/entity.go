package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Start(*StateShape) error
	Update(*StateShape) error
	Draw(*StateShape, *ebiten.Image) error
}
