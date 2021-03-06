package main

import (
	"fmt"

	"github.com/slowfrog/stephen/pkg/display"
	"github.com/slowfrog/stephen/pkg/model"
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
	fmt.Printf("%#v\n", b)

	s := make([]model.Sausage, 2)
	s[0].Alignment = model.HORIZONTAL
	s[1].Pos.X = 3
	s[1].Pos.Y = 1

	st := model.Stephen{model.Pos{4, 2}, model.LEFT}
	w := model.NewWorld(b, s, st)

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

	running = true
	for running {
		moved := false
		var d model.Dir
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				kev := event.(*sdl.KeyDownEvent)
				switch (*kev).Keysym.Sym {
				case sdl.K_ESCAPE:
					running = false
				case sdl.K_UP:
					d = model.UP
					moved = true
				case sdl.K_DOWN:
					d = model.DOWN
					moved = true
				case sdl.K_LEFT:
					d = model.LEFT
					moved = true
				case sdl.K_RIGHT:
					d = model.RIGHT
					moved = true
				}
			}
		}

		if moved {
			w.MoveStephen(d)
		}

		ds.Clear(GRAY)
		renderWorld(ds, w)
		ds.Present()
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

var TRANS_ORANGE = display.Rgba(255, 128, 0, 128)
var TRANS_RED = display.Rgba(255, 0, 0, 128)

const (
	OX = 10
	OY = 10
	CW = 50
	CH = 50
)

func renderWorld(ds *display.State, w *model.World) {
	renderBoard(ds, w.Board())
	for _, s := range w.Sausages() {
		renderSausage(ds, &s)
	}
	renderStephen(ds, w.Stephen())
}

func renderBoard(ds *display.State, b *model.Board) {
	w := int8(b.Width())
	h := int8(b.Height())
	rect := sdl.Rect{W: CW, H: CH}
	var x, y int8
	for x = 0; x < w; x++ {
		rect.X = int32(x)*CW + OX
		for y = 0; y < h; y++ {
			rect.Y = int32(y)*CH + OY
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

const (
	SOX = 5
	SOY = 5
	SW  = CW - 2*SOX
	SH  = CH - 2*SOY
)

func renderSausage(ds *display.State, s *model.Sausage) {
	var sw, sh int32
	if s.Alignment == model.HORIZONTAL {
		sw, sh = (CW-SOX)*2, CH-SOY*2
	} else {
		sw, sh = CW-SOX*2, (CH-SOY)*2
	}
	p := s.GetPos()
	rect := sdl.Rect{X: int32(p.X)*CW + OX + SOX, Y: int32(p.Y)*CH + OY + SOY, W: sw, H: sh}
	ds.FillRect(&rect, TRANS_RED)
}

const (
	STOX = 10
	STOY = 10
	STOIX = 20
	STOIY = 20
	STW  = CW - 2*STOX
	STH  = CH - 2*STOY
	STIW  = CW - 2*STOIX
	STIH  = CH - 2*STOIY
)

func min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func renderStephen(ds *display.State, s *model.Stephen) {
	var rects [2]sdl.Rect
	rects[0].X = int32(s.Pos.X) * CW + OX + STOX
	rects[0].Y = int32(s.Pos.Y)*CH + OY + STOY
	rects[0].W = STW
	rects[0].H = STH
	dx, dy := s.Dir.Offset()
	minX := min(int32(s.Pos.X), int32(s.Pos.X + dx))
	minY := min(int32(s.Pos.Y), int32(s.Pos.Y + dy))
	adX := abs(int32(dx))
	adY := abs(int32(dy))
	
	rects[1].X = int32(minX) * CW + OX + STOIX
	rects[1].Y = int32(minY) * CH + OY + STOIY
	rects[1].W = STIW + int32(adX * CW)
	rects[1].H = STIH + int32(adY * CH)

	ds.FillRects(rects[0:2], YELLOW)
}
