package enemy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pilotpirxie/ebiten-test/src/game"
	"os"
	"path"
)

type Enemy struct {
	enemyImage *ebiten.Image
	x          float64
	y          float64
	speed      float64
}

func NewEnemy(x float64, y float64, speed float64) game.Entity {
	return &Enemy{
		x:     x,
		y:     y,
		speed: speed,
	}
}

func (e *Enemy) Start(_ *game.StateShape) error {
	var err error
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	e.enemyImage, _, err = ebitenutil.NewImageFromFile(path.Join(pwd, "../src/assets/textures/enemy.png"))
	if err != nil {
		return err
	}

	return nil
}

func (e *Enemy) Update(_ *game.StateShape) error {
	return nil
}

func (e *Enemy) Draw(_ *game.StateShape, image *ebiten.Image) error {
	if e.enemyImage == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.x, e.y)
	image.DrawImage(e.enemyImage, op)

	return nil
}
