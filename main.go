package main

import (
	"fmt"

	"github.com/slowfrog/stephen/display"
	"github.com/slowfrog/stephen/model"
	"github.com/veandco/go-sdl2/sdl"
)

const W int32 = 400
const H int32 = 400

func main() {
	var event sdl.Event
	var err error
	var running bool

	b := model.CreateBoard(5, 5)
	fmt.Printf("Board: %s\n", b)
	/*err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("The Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)*/
	ds, err := display.InitState(W, H)
	if err != nil {
		panic(err)
	}
	defer ds.Destroy()
	
	points := make([]sdl.Point, 5)
	points[0] = sdl.Point{10, 10}
	points[1] = sdl.Point{100, 10}
	points[2] = sdl.Point{100, 200}
	points[3] = sdl.Point{150, 150}
	points[4] = sdl.Point{10, 10}

	var dx, dy int32
	dx, dy = 1, 1

	renderFigure(ds.Renderer(), points, &dx, &dy)
	
	//window.UpdateSurface()

	running = true
	for running {
		renderFigure(ds.Renderer(), points, &dx, &dy)
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				kev := event.(*sdl.KeyDownEvent)
				if (*kev).Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
			default:
				fmt.Printf("Ev %s\n", t)
			}
		}
		sdl.Delay(5)
	}
}

func renderFigure(renderer *sdl.Renderer, points []sdl.Point, dx *int32, dy *int32) {
	err := renderer.SetDrawColor(0, 0, 0, 255)
	if err != nil {
		panic(err)
	}

	renderer.Clear()
	
	err = renderer.SetDrawColor(255, 128, 0, 255)
	if err != nil {
		panic(err)
	}

	points[2].X += *dx
	points[2].Y += *dy

	if points[2].X >= W {
		*dx = -1
	} else if points[2].X <= 0 {
		*dx = 1
	}
	if points[2].Y >= H {
		*dy = -1
	} else if points[2].Y <= 0 {
		*dy = 1
	}
	
	err = renderer.DrawLines(points)
	if err != nil {
		panic(err)
	}

	renderer.Present()
}
