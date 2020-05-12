package main

import (
	"golang.org/x/image/colornames"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title: "Game Of Life",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	
	imd.Color = colornames.Black
	imd.Push(pixel.V(100, 100))
	imd.Push(pixel.V(200, 200))
	imd.Rectangle(0)

	for !win.Closed() {
		win.Clear(colornames.White)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
