package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 856
	screenHeight = 480
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("hero.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	player struct {
		heroX int
		heroY int
	}
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.heroX = g.player.heroX - 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.heroX = g.player.heroX + 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.heroY = g.player.heroY - 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.heroY = g.player.heroY + 1
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.player.heroX), float64(g.player.heroY))
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten Demo")
	if err := ebiten.RunGame(&Game{
		player: struct {
			heroX int
			heroY int
		}{heroX: 0, heroY: 0},
	}); err != nil {
		log.Fatal(err)
	}
}
