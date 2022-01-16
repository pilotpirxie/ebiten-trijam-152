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
	x         float64
	y         float64
	speed     float64
}

func NewPlayer(x float64, y float64, speed float64) game.Entity {
	return &Player{
		x:     x,
		y:     y,
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
		p.x = p.x - p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.x = p.x + p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.y = p.y - p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.y = p.y + p.speed
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
	op.GeoM.Translate(p.x, p.y)
	image.DrawImage(p.heroImage, op)

	return nil
}