package main

import (
	"fmt"

	// "github.com/veandco/go-sdl2/mix"
	sdl "github.com/veandco/go-sdl2/sdl"
	// "github.com/veandco/go-sdl2/ttf"
	mog "simple-sdl2-project/internal/mog"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(mog.MW, "\nFailed to initialize a window: %v", err)
		panic("Failed to initialize a window")
	}

	window, err := sdl.CreateWindow("Simple SDL2 Project",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)

	if err != nil {
		fmt.Fprintf(mog.MW, "\nFailed to create the SDL Window: %v", err)
		panic("Failed to create the SDL Window")
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(mog.MW, "\nFailed to Initialize renderer: %v", err)
		panic("Failed to Initialize renderer")
	}
	defer renderer.Destroy()
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				sdl.Quit()
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		updateAndDraw(renderer)

		// th.Do(func() {
		// 	fmt.Printf("\n(%d, %d, %d)", mouseX, mouseY, mouseButtonDown)
		// })

		renderer.Present()
	}
}

func updateAndDraw(renderer *sdl.Renderer) {
	// nothing
}
