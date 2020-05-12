package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	cells          = 250
	cellWidth      = 4
	windowSize     = float64(cells * cellWidth)
	initalCellsPct = 0.3
)

type gameOfLife struct {
	currentState [][]bool
	nextState    [][]bool
	size         int
}

func (g *gameOfLife) initialize() {
	g.currentState = make([][]bool, cells)
	g.nextState = make([][]bool, cells)
	for i := range g.currentState {
		g.currentState[i] = make([]bool, cells)
		g.nextState[i] = make([]bool, cells)
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

func (g *gameOfLife) checkNeighbors(x, y int) int {
	n := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			posX, posY := i, j
			// fmt.Println(posX, posY)
			if posX == x && posY == y {
				continue
			}
			if posX < 0 || posX >= cells {
				continue
			}
			if posY < 0 || posY >= cells {
				continue
			}
			if g.currentState[posX][posY] {
				n++
			}
		}
	}

	return n
}

func (g *gameOfLife) calculateNextState() {
	for i, row := range g.currentState {
		for j, v := range row {
			n := g.checkNeighbors(i, j)
			g.nextState[i][j] = false
			if v {
				if n == 2 || n == 3 {
					g.nextState[i][j] = true
				}
			} else {
				if n == 3 {
					g.nextState[i][j] = true
				}
			}
		}
	}
	g.currentState, g.nextState = g.nextState, g.currentState
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
	game.initialize()

	imd := *imdraw.New(nil)

	for !win.Closed() {
		win.Clear(colornames.White)
		game.draw(&imd)
		imd.Draw(win)
		game.calculateNextState()
		win.Update()
		time.Sleep(40 * time.Millisecond)
	}
}

func main() {
	pixelgl.Run(run)
}
