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

	b := model.CreateBoard(6, 4)
	b.Set(1, 0, model.GROUND).Set(3, 2, model.GRILL).Set(5, 3, model.GROUND)
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

	renderFigure(ds, points, &dx, &dy)
	
	//window.UpdateSurface()
	
	running = true
	for running {
		ds.Clear(GRAY)
		renderBoard(ds, &b)
		renderFigure(ds, points, &dx, &dy)
		ds.Present()
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

var BLACK = display.Rgb(0, 0, 0)
var GRAY = display.Rgb(128, 128, 128)
var ORANGE = display.Rgb(255, 128, 0)
var GREEN = display.Rgb(0, 255, 0)
var BLUE = display.Rgb(0, 0, 255)
var YELLOW = display.Rgb(255, 255, 0)
var RED = display.Rgb(255, 0, 0)

func renderBoard(ds *display.State, b *model.Board) {
	w := int8(b.Width())
	h := int8(b.Height())
	rect := sdl.Rect{W:50, H:50}
	var x, y int8
	for x = 0; x < w; x++ {
		rect.X = int32(x) * 50 + 10
		for y = 0; y < h; y++ {
			rect.Y = int32(y) * 50 + 10
			c := b.Get(x, y)
			col := BLACK
			if c == model.GROUND {
				col = GREEN
			} else if c == model.GRILL {
				col = ORANGE
			}
			err := ds.FillRect(&rect, col)
			if err != nil {
				panic(err)
			}
		}
	}
}

func renderFigure(ds *display.State, points []sdl.Point, dx *int32, dy *int32) {
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
	
	err := ds.DrawLines(points, YELLOW)
	if err != nil {
		panic(err)
	}

	r := sdl.Rect{points[2].X - 10, points[2].Y - 10, 20, 20}
	err = ds.FillRect(&r, RED)
	if err != nil {
		panic(err)
	}
}
