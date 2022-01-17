package controller

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pilotpirxie/ebiten-trijam-152/src/game"
	"github.com/pilotpirxie/ebiten-trijam-152/src/prefabs/audio"
	"time"
)

type Controller struct {
	previousScore int
}

func NewController() game.Entity {
	return &Controller{}
}

func (b *Controller) Start(_ *game.StateShape) error {
	return nil
}

func (b *Controller) Update(g *game.StateShape) error {
	if !g.GameOver {
		if g.Score != b.previousScore && g.Score%2 == 0 {
			err, entity := g.GetEntity("audio", &audio.Audio{})
			if err != nil {
				panic(nil)
			}

			audioEntity := entity.(*audio.Audio)
			err = audioEntity.PlayCoin()
			if err != nil {
				panic(nil)
			}

			b.previousScore = g.Score
		}

		g.Score = int(time.Now().Sub(g.StartTime).Seconds())
	}

	return nil
}

func (b *Controller) Draw(g *game.StateShape, screen *ebiten.Image) error {
	if !g.GameOver {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %ds", g.Score))
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("GAME OVER! Score: %ds", g.Score))
	}

	return nil
}
