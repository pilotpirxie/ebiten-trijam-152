package enemy

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pilotpirxie/ebiten-test/src/game"
	player2 "github.com/pilotpirxie/ebiten-test/src/prefabs/player"
	"math"
	"os"
	"path"
)

type Enemy struct {
	enemyImage *ebiten.Image
	X          float64
	Y          float64
	speed      float64
}

func NewEnemy(x float64, y float64, speed float64) game.Entity {
	return &Enemy{
		X:     x,
		Y:     y,
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

func (e *Enemy) Update(g *game.StateShape) error {
	err, entity := g.GetEntity("player", &player2.Player{})
	if err != nil {
		return nil
	}

	player := entity.(*player2.Player)

	distance := math.Sqrt(math.Pow(player.X-e.X, 2) + math.Pow(player.Y-e.Y, 2))
	if distance > 10 {
		fmt.Println("elo")
	}
	return nil
}

func (e *Enemy) Draw(_ *game.StateShape, image *ebiten.Image) error {
	if e.enemyImage == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.X, e.Y)
	image.DrawImage(e.enemyImage, op)

	return nil
}
