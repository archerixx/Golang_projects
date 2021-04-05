package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWSPOS_UNDEFINED,
		int32(800), int32(600), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
	}

	defer window.Destory()
}
