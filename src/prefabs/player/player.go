package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pilotpirxie/ebiten-test/src/game"
	"os"
	"path"
)

type Player struct {
	walkSpeed float64
	runSpeed  float64
	image     *ebiten.Image
	X         float64
	Y         float64
	speed     float64
	isFlipped bool
}

func NewPlayer(x float64, y float64, speed float64) game.Entity {
	return &Player{
		X:     x,
		Y:     y,
		speed: speed,
	}
}

func (p *Player) Start(_ *game.StateShape) error {
	p.walkSpeed = 3
	p.runSpeed = 4
	p.X = game.ScreenWidth / 2
	p.Y = game.ScreenHeight / 2

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	p.image, _, err = ebitenutil.NewImageFromFile(path.Join(pwd, "../src/assets/textures/hero.png"))
	if err != nil {
		return err
	}

	return nil
}

func (p *Player) Update(g *game.StateShape) error {
	if g.GameOver {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.X = p.X - p.speed
		p.isFlipped = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.X = p.X + p.speed
		p.isFlipped = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y = p.Y - p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y = p.Y + p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
		p.speed = p.runSpeed
	} else {
		p.speed = p.walkSpeed
	}

	return nil
}

func (p *Player) Draw(_ *game.StateShape, image *ebiten.Image) error {
	if p.image == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}

	if p.isFlipped {
		op.GeoM.Scale(-0.3, 0.3)
	} else {
		op.GeoM.Scale(0.3, 0.3)
	}

	op.GeoM.Translate(p.X, p.Y)
	image.DrawImage(p.image, op)

	return nil
}
