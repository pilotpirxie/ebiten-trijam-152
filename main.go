package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 854
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

type Player struct {
	heroX float64
	heroY float64
	speed float64
}

type Game struct {
	player Player
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.heroX = g.player.heroX - g.player.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.heroX = g.player.heroX + g.player.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.heroY = g.player.heroY - g.player.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.heroY = g.player.heroY + g.player.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
		g.player.speed = 2
	} else {
		g.player.speed = 1
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.NRGBA{
		R: 25,
		G: 25,
		B: 25,
		A: 255,
	})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.player.heroX), float64(g.player.heroY))
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten Demo")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(&Game{
		player: Player{heroX: 0, heroY: 0, speed: 1},
	}); err != nil {
		log.Fatal(err)
	}
}
