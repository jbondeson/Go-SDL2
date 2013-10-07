package main

import (
	"fmt"
	"github.com/krig/Go-SDL2/sdl"
)

func loadImage(name string) *sdl.Surface {
	image := sdl.Load(name)

	if image == nil {
		panic(sdl.GetError())
	}

	return image

}

func main() {
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		panic(sdl.GetError())
	}

	defer sdl.Quit()

	window, rend := sdl.CreateWindowAndRenderer(640, 480, sdl.WINDOW_SHOWN |
		sdl.RENDERER_ACCELERATED |
		sdl.RENDERER_PRESENTVSYNC)

	if (window == nil) || (rend == nil) {
		fmt.Printf("%#v\n", sdl.GetError())
	}

	window.SetTitle("Podcast Studio")

	running := true
	event := &sdl.Event{}
	for running {
		for event.Poll() {
			switch event.Type {
			case sdl.QUIT:
				running = false
			}
		}
		rend.SetDrawColor(sdl.Color{0x30, 0xff, 0x30, 0xFF, 0x00})
		//rect := &sdl.Rect{0, 0, (uint16)(window.W), (uint16)(window.H)}
		//rend.FillRect(rect)
		rend.FillRect(nil)
		rend.Present()
	}
}
