package bg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pilotpirxie/ebiten-test/src/game"
	"os"
	"path"
)

type Background struct {
	image *ebiten.Image
}

func NewBackground(x float64, y float64, speed float64) game.Entity {
	return &Background{}
}

func (b *Background) Start(_ *game.StateShape) error {
	var err error
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	b.image, _, err = ebitenutil.NewImageFromFile(path.Join(pwd, "../src/assets/textures/bg.png"))
	if err != nil {
		return err
	}

	return nil
}

func (b *Background) Update(g *game.StateShape) error {
	return nil
}

func (b *Background) Draw(_ *game.StateShape, image *ebiten.Image) error {
	if b.image == nil {
		return nil
	}

	image.DrawImage(b.image, &ebiten.DrawImageOptions{})

	return nil
}