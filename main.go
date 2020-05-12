package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	cells          = 40
	cellWidth      = 20
	windowSize     = float64(cells * cellWidth)
	initalCellsPct = 0.3
)

type gameOfLife struct {
	currentState [][]bool
	size         int
}

func (g *gameOfLife) initialize() {
	g.currentState = make([][]bool, cells)
	for i := range g.currentState {
		g.currentState[i] = make([]bool, cells)
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			r := rand.Float32()
			if r <= initalCellsPct {
				g.currentState[i][j] = true
			} else {
				g.currentState[i][j] = false
			}
		}
	}
}

func (g *gameOfLife) draw(imd *imdraw.IMDraw) {
	imd.Clear()
	for i, row := range g.currentState {
		for j, v := range row {
			if v {
				x0, y0 := float64(i*cellWidth), float64(j*cellWidth)
				x1, y1 := float64(i*cellWidth+cellWidth), float64(j*cellWidth+cellWidth)

				imd.Color = colornames.Black
				imd.Push(pixel.V(x0, y0))
				imd.Push(pixel.V(x1, y1))
				imd.Rectangle(0)
			}
		}
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game Of Life",
		Bounds: pixel.R(0, 0, windowSize, windowSize),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	game := gameOfLife{}
	// game.currentState = make([][]bool, cells)
	game.initialize()

	imd := *imdraw.New(nil)

	for !win.Closed() {
		win.Clear(colornames.White)
		game.draw(&imd)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
