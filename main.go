package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const (
	cells          = 200
	cellWidth      = 5
	gameSize       = float64(cells * cellWidth)
	initalCellsPct = 0.3
	menuWidth      = 400
	fps            = 15
)

var (
	paused = false
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

			if posX == -1 {
				posX = cells - 1
			} else if posX == cells {
				posX = 0
			}

			if posY == -1 {
				posY = cells - 1
			} else if posY == cells {
				posY = 0
			}

			if posX == x && posY == y {
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

func drawMenu(imd *imdraw.IMDraw, w *pixelgl.Window) {

	imd.Color = color.RGBA{0xb2, 0xeb, 0xf2, 0xff}
	imd.Push(pixel.V(float64(gameSize), float64(0)))
	imd.Push(pixel.V(float64(gameSize+400), float64(gameSize)))
	imd.Rectangle(0)

	imd.Draw(w)

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	pauseText := text.New(pixel.V(1050, 900), atlas)
	pauseText.Color = colornames.Black
	pauseText.LineHeight = atlas.LineHeight() * 1.5

	var text string
	if paused {
		text = "Press P to resume"
	} else {
		text = "Press P to pause"
	}

	textScale := float64(2.3)
	pauseText.Orig.X = gameSize + menuWidth/2 - pauseText.BoundsOf(text).W()/2*textScale
	pauseText.Dot.X = gameSize + menuWidth/2 - pauseText.BoundsOf(text).W()/2*textScale
	fmt.Fprintln(pauseText, text)

	pauseText.Draw(w, pixel.IM.Scaled(pauseText.Orig, textScale))
}

func handlePause(w *pixelgl.Window) {
	if w.JustPressed(pixelgl.KeyP) {
		paused = !paused
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game Of Life",
		Bounds: pixel.R(0, 0, gameSize+menuWidth, gameSize),
		// VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	game := gameOfLife{}
	game.initialize()

	imd := *imdraw.New(nil)

	f := time.Tick(time.Second / fps)

	for !win.Closed() {
		handlePause(win)

		if !paused {
			win.Clear(colornames.White)
			game.draw(&imd)
			imd.Draw(win)
			game.calculateNextState()
		}
		drawMenu(&imd, win)

		win.Update()

		<-f
	}
}

func main() {
	pixelgl.Run(run)
}
