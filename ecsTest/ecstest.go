// https://threedots.tech/post/making-games-in-go-ecs/
package ecstest

import (
	"fmt"
	"image/color"
	"io"

	"github.com/yohamta/donburi"
)

type PositionData struct {
	X float64
	Y float64
}

type ColorData struct {
	Color color.Color
}

type Tile struct {
	Position PositionData
	Color    ColorData
}

type Game struct {
}

func (g *Game) Update() error {
	return nil
}
func (g *Game) Draw()

var (
	Position = donburi.NewComponentType[PositionData]()
	Color    = donburi.NewComponentType[ColorData]()
)

func Register(registrar map[string]func(io.Writer)) map[string]func(io.Writer) {
	registrar["ecsTest"] = Run
	return registrar
}

func Run(writer io.Writer) {
	world := donburi.NewWorld()
	entity := world.Create(Position, Color)

	fmt.Fprintln(writer, entity.Version())

	entry := world.Entry(entity)

	fmt.Fprintln(writer, Position.Get(entry))
	fmt.Fprintln(writer, Color.Get(entry))
}
