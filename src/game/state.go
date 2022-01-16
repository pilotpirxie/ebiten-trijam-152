package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"reflect"
	"time"
)

const (
	ScreenWidth  = 854
	ScreenHeight = 480
)

type StateShape struct {
	StartTime time.Time
	Entities  []Entity
	Globals   map[string]interface{}
	Score     int
	GameOver  bool
}

var State *StateShape

func init() {
	State = &StateShape{
		StartTime: time.Now(),
		Entities:  []Entity{},
		Globals:   map[string]interface{}{},
		Score:     0,
	}
}

func (g *StateShape) Update() error {
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

	if !g.GameOver {
		g.Score = int(time.Now().Sub(g.StartTime).Seconds())
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %ds", g.Score))
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("GAME OVER! Score: %ds", g.Score))
	}
}

func (g *StateShape) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *StateShape) GetEntity(entityName string, entityType interface{}) (error, interface{}) {
	if !g.EntityExists(entityName) {
		return fmt.Errorf("entity with specified name doesn't exist"), nil
	}

	entity := g.Globals[entityName]

	if reflect.TypeOf(entity) == reflect.TypeOf(entityType) {
		return nil, entity
	} else {
		return fmt.Errorf("invalid type"), nil
	}
}

func (g *StateShape) EntityExists(entityName string) bool {
	entity := g.Globals[entityName]

	if entity == nil {
		return false
	} else {
		return true
	}
}

func (g *StateShape) AddEntity(entityName string, entity Entity) error {
	if g.EntityExists(entityName) {
		return fmt.Errorf("entity with specified name already exists")
	}

	g.Globals[entityName] = entity
	g.Entities = append(g.Entities, entity)

	err := entity.Start(g)
	if err != nil {
		return err
	}

	return nil
}

func (g *StateShape) DestroyEntity(entityName string) error {
	if !g.EntityExists(entityName) {
		return fmt.Errorf("entity with specified name doesn't exist")
	}

	entity := g.Globals[entityName]

	for i := range g.Entities {
		if g.Entities[i] == entity {
			g.Entities = append(g.Entities[:i], g.Entities[i+1:]...)
		}
	}

	delete(g.Globals, entityName)

	return nil
}
