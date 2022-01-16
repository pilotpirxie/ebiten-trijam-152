package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pilotpirxie/ebiten-test/src/game"
	"github.com/pilotpirxie/ebiten-test/src/prefabs/enemy"
	"github.com/pilotpirxie/ebiten-test/src/prefabs/player"
	_ "image/png"
	"log"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Ebiten Demo")
	ebiten.SetWindowResizable(true)

	game.State.Entities = append(
		game.State.Entities,
		&player.Player{},
		&enemy.Enemy{},
	)

	if err := ebiten.RunGame(game.State); err != nil {
		log.Fatal(err)
	}
}
