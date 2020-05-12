package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	cells      = 40
	cellWidth  = 20
	windowSize = float64(cells * cellWidth)
)

// type GameOfLife struct {
// 	size int
// }

func draw(imd *imdraw.IMDraw) {
	imd.Clear()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if (i%2 == 0 && j%2 == 1) || (i%2 == 1 && j%2 == 0) {
				imd.Color = colornames.White
			} else {
				imd.Color = colornames.Black
			}

			imd.Push(pixel.V(float64(i*cellWidth), float64(j*cellWidth)))
			imd.Push(pixel.V(float64(i*cellWidth+cellWidth), float64(j*cellWidth+cellWidth)))
			imd.Rectangle(0)
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

	imd := *imdraw.New(nil)

	for !win.Closed() {
		win.Clear(colornames.White)
		draw(&imd)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
