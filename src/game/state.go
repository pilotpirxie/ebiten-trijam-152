package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"time"
)

const (
	ScreenWidth  = 854
	ScreenHeight = 480
)

type StateShape struct {
	StartTime time.Time
	Entities  []Entity
	Score     int
}

var State *StateShape
var started bool

func init() {
	State = &StateShape{
		StartTime: time.Now(),
		Entities:  []Entity{},
		Score:     0,
	}
}

func (g *StateShape) Update() error {
	if !started {
		for _, entity := range State.Entities {
			err := entity.Start(State)
			if err != nil {
				return err
			}
		}

		started = true
	}

	for _, entity := range g.Entities {
		err := entity.Update(g)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *StateShape) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{
		R: 25,
		G: 25,
		B: 25,
		A: 255,
	})

	for _, entity := range g.Entities {
		err := entity.Draw(g, screen)
		if err != nil {
			panic(err)
		}
	}

	g.Score = int(time.Now().Sub(g.StartTime).Seconds() * 10)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.Score))
}

func (g *StateShape) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}
