package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pilotpirxie/ebiten-test/src/data"
	"github.com/pilotpirxie/ebiten-test/src/game"
	_ "image/png"
	"log"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Ebiten Demo")
	ebiten.SetWindowResizable(true)

	err := game.State.AddEntity("background", &data.Background)
	if err != nil {
		log.Fatal(err)
	}

	err = game.State.AddEntity("player", &data.Player)
	if err != nil {
		log.Fatal(err)
	}

	err = game.State.AddEntity("enemy", &data.Enemy)
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game.State); err != nil {
		log.Fatal(err)
	}
}
