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

	//game.State.Entities = append(
	//	game.State.Entities,
	//	&data.Enemy,
	//	&data.Player,
	//)
	//
	//game.State.Globals["player"] = &data.Player
	//game.State.Globals["enemy"] = &data.Player

	game.State.AddEntity("player", &data.Player)
	game.State.AddEntity("enemy", &data.Enemy)

	if err := ebiten.RunGame(game.State); err != nil {
		log.Fatal(err)
	}
}
