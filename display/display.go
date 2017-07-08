package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

func InitSdl() (window *sdl.Window, err error) {
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, err
	}

	window, err = sdl.CreateWindow("The Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	return window, err
}

