package main

import (
	"fmt"
	"github.com/krig/Go-SDL2/mixer"
	"github.com/krig/Go-SDL2/sdl"
	"github.com/krig/Go-SDL2/ttf"
	"log"
	"math"
)

type Point struct {
	x int
	y int
}

func (a Point) add(b Point) Point { return Point{a.x + b.x, a.y + b.y} }

func (a Point) sub(b Point) Point { return Point{a.x - b.x, a.y - b.y} }

func (a Point) length() float64 { return math.Sqrt(float64(a.x*a.x + a.y*a.y)) }

func (a Point) mul(b float64) Point {
	return Point{int(float64(a.x) * b), int(float64(a.y) * b)}
}

func worm(in <-chan Point, out chan<- Point, draw chan<- Point) {

	t := Point{0, 0}

	for {
		p := (<-in).sub(t)

		if p.length() > 48 {
			t = t.add(p.mul(0.1))
		}

		draw <- t
		out <- t
	}
}

func RenderTextToTexture(r *sdl.Renderer, f *ttf.Font, text string, color sdl.Color) (*sdl.Texture, int, int) {
	textw, texth, err := f.SizeText(text)
	if err != nil {
		log.Fatal(err)
	}
	txt_surface := f.RenderText_Blended(text, color)
	txt_tex := r.CreateTextureFromSurface(txt_surface)
	txt_surface.Free()
	return txt_tex, textw, texth
}

func main() {
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		log.Fatal(sdl.GetError())
	}
	defer sdl.Quit()

	if mixer.OpenAudio(mixer.DEFAULT_FREQUENCY, mixer.DEFAULT_FORMAT,
		mixer.DEFAULT_CHANNELS, 4096) != 0 {
		log.Fatal(sdl.GetError())
	}

	window, rend := sdl.CreateWindowAndRenderer(640, 480, sdl.WINDOW_SHOWN)

	if window == nil {
		log.Println("nil window")
		log.Fatal(sdl.GetError())
	}
	defer window.Destroy()

	if rend == nil {
		log.Println("nil rend")
		log.Fatal(sdl.GetError())
	}
	defer rend.Destroy()

	if ttf.Init() != 0 {
		log.Fatal(sdl.GetError())
	}
	defer ttf.Quit()

	window.SetTitle("First SDL2 Window")

	image := sdl.Load("./test.png")
	defer image.Free()

	if image == nil {
		log.Println("nil image")
		log.Fatal(sdl.GetError())
	}

	window.SetIcon(image)

	tex := rend.CreateTextureFromSurface(image)
	defer tex.Destroy()

	font := ttf.OpenFont("./Fontin Sans.otf", 16)
	defer font.Close()
	txt_tex, _, _ := RenderTextToTexture(rend, font, "This is a test", sdl.Color{0x7F, 0xFF, 0x10, 0xFF})
	defer txt_tex.Destroy()

	running := true

	worm_in := make(chan Point)
	draw := make(chan Point, 64)

	var out chan Point
	var in chan Point

	out = worm_in

	in = out
	out = make(chan Point)
	go worm(in, out, draw)

	// ticker := time.NewTicker(time.Second / 50) // 50 Hz

	window.ShowSimpleMessageBox(sdl.MESSAGEBOX_INFORMATION, "Test Message", "SDL2 supports message boxes!")
	event := &sdl.Event{}

	for running {
//		select {
		/*
					case <-ticker.C:
						rend.SetDrawColor(0x30, 0x20, 0x19, 0xFF)
						rend.FillRect(nil)
			            if sdl.GetError() != "" {
			                log.Fatalf(sdl.GetError())
			            }

					loop:
						for {
							select {
							case p := <-draw:
								rend.Clear()
								rend.Copy(tex, &sdl.Rect{int16(p.x), int16(p.y), 0, 0}, nil)

							case <-out:
							default:
								break loop
							}
						}

						var p Point
						sdl.GetMouseState(&p.x, &p.y)
						worm_in <- p

						rend.Present()
		*/
		for event.Poll() {
			switch e := event.Get().(type) {
			case sdl.QuitEvent:
				running = false

			case sdl.KeyboardEvent:
				println("")
				println(e.Keysym.Keycode, ": ", sdl.GetKeyName(e.Keysym.Keycode))

				if e.Keysym.Keycode == sdl.K_ESCAPE {
					running = false
				}

				fmt.Printf("%04x ", e.Type)

				println()

				fmt.Printf("Type: %02x State: %02x\n", e.Type, e.State)
				fmt.Printf("Scancode: %02x Keycode: %02x Mod: %04x\n",
					e.Keysym.Scancode,
					e.Keysym.Keycode,
					e.Keysym.Mod)
			case sdl.MouseButtonEvent:
				if e.Type == sdl.MOUSEBUTTONDOWN {
					println("Click:", e.X, e.Y)
					in = out
					out = make(chan Point)
					go worm(in, out, draw)
				}
			}
		}
		rend.SetDrawColor(sdl.Color{0x30, 0x20, 0x19, 0xFF})
		rend.FillRect(nil)
		rend.Copy(tex, nil, nil)

		rend.Copy(txt_tex, nil, nil)

		rend.Present()
	}
}
