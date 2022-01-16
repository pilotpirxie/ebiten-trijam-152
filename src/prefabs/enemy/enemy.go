package enemy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pilotpirxie/ebiten-test/src/game"
	"github.com/pilotpirxie/ebiten-test/src/prefabs/player"
	"math"
	"os"
	"path"
)

type Enemy struct {
	image     *ebiten.Image
	X         float64
	Y         float64
	speed     float64
	isFlipped bool
}

func NewEnemy(x float64, y float64, speed float64) game.Entity {
	return &Enemy{
		X:     x,
		Y:     y,
		speed: speed,
	}
}

func (e *Enemy) Start(_ *game.StateShape) error {
	e.speed = 2

	var err error
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	e.image, _, err = ebitenutil.NewImageFromFile(path.Join(pwd, "../src/assets/textures/enemy.png"))
	if err != nil {
		return err
	}

	return nil
}

func (e *Enemy) Update(g *game.StateShape) error {
	err, entity := g.GetEntity("player", &player.Player{})
	if err != nil {
		return nil
	}

	playerEntity := entity.(*player.Player)

	distance := math.Sqrt(math.Pow(playerEntity.X-e.X, 2) + math.Pow(playerEntity.Y-e.Y, 2))
	if distance > 20 {
		if playerEntity.X > e.X {
			e.X += e.speed
			e.isFlipped = false
		} else {
			e.X -= e.speed
			e.isFlipped = true
		}

		if playerEntity.Y > e.Y {
			e.Y += e.speed
		} else {
			e.Y -= e.speed
		}
	} else {
		g.GameOver = true
	}
	return nil
}

func (e *Enemy) Draw(_ *game.StateShape, image *ebiten.Image) error {
	if e.image == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}

	if e.isFlipped {
		op.GeoM.Scale(-0.3, 0.3)
	} else {
		op.GeoM.Scale(0.3, 0.3)
	}

	op.GeoM.Translate(e.X, e.Y)
	image.DrawImage(e.image, op)

	return nil
}
