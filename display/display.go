package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func Rgb(r, g, b uint8) (c Color) {
	c.R = r
	c.G = g
	c.B = b
	c.A = 255
	return
}

func Rgba(r, g, b, a uint8) (c Color) {
	c.R = r
	c.G = g
	c.B = b
	c.A = a
	return
}

// State captures the whole display state in an opaque way
type State struct {
	w *sdl.Window
	r *sdl.Renderer
}

// InitState create a display state and opens the main window, with a renderer
func InitState(width, height int32) (s *State, err error) {
	// SDL global init
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return
	}

	// Main window
	s = new(State)
	s.w, err = sdl.CreateWindow("The Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int(width), int(height), sdl.WINDOW_SHOWN)
	if err != nil {
		s.Destroy()
		return
	}

	// Renderer associated with the main window
	s.r, err = sdl.CreateRenderer(s.w, -1, 0)
	if err != nil {
		s.Destroy()
		return
	}
	return
}

// Destroy closes everything that had been initialized
func (s *State) Destroy() {
	if s.r != nil {
		s.r.Destroy()
		s.r = nil
	}
	if s.w != nil {
		s.w.Destroy()
		s.w = nil
	}
	sdl.Quit()
}

// Window gives access to the main SDL window
func (s *State) Window() *sdl.Window {
	return s.w
}

// Renderer gives access to the renderer of the main SDL window
func (s *State) Renderer() *sdl.Renderer {
	return s.r
}

func (s *State) setColor(c Color) error {
	return s.r.SetDrawColor(c.R, c.G, c.B, c.A)
}

func (s *State) Clear(c Color) (err error) {
	err = s.setColor(c)
	if err != nil {
		return err
	}
	err = s.r.Clear()
	return
}

func (s *State) FillRect(rect *sdl.Rect, c Color) (err error) {
	err = s.setColor(c)
	if err != nil {
		return
	}
	err = s.r.FillRect(rect)
	return
}

func (s *State) FillRects(rects []sdl.Rect, c Color) (err error) {
	err = s.setColor(c)
	if err != nil {
		return
	}
	err = s.r.FillRects(rects)
	return
}

func (s *State) DrawLines(points []sdl.Point, c Color) (err error) {
	err = s.setColor(c)
	if err != nil {
		return
	}
	err = s.r.DrawLines(points)
	return
}

func (s *State) Present(){
	s.r.Present()
}

