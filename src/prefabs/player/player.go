package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pilotpirxie/ebiten-test/src/game"
	"os"
	"path"
)

const walkSpeed = 5
const runSpeed = 8

type Player struct {
	heroImage *ebiten.Image
	X         float64
	Y         float64
	speed     float64
}

func NewPlayer(x float64, y float64, speed float64) game.Entity {
	return &Player{
		X:     x,
		Y:     y,
		speed: speed,
	}
}

func (p *Player) Start(_ *game.StateShape) error {
	var err error
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	p.heroImage, _, err = ebitenutil.NewImageFromFile(path.Join(pwd, "../src/assets/textures/hero.png"))
	if err != nil {
		return err
	}

	return nil
}

func (p *Player) Update(_ *game.StateShape) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.X = p.X - p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.X = p.X + p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y = p.Y - p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y = p.Y + p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
		p.speed = runSpeed
	} else {
		p.speed = walkSpeed
	}

	return nil
}

func (p *Player) Draw(_ *game.StateShape, image *ebiten.Image) error {
	if p.heroImage == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	image.DrawImage(p.heroImage, op)

	return nil
}
