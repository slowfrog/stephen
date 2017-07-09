package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

// State captures the whole display state in an opaque way
type State struct {
	w *sdl.Window
	r *sdl.Renderer
}

// InitState create a display state and opens the main window, with a renderer
func InitState(width, height int32) (s State, err error) {
	// SDL global init
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return
	}

	// Main window
	s = State{}
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
