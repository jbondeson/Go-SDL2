package sdl

/*
  SDL Go Wrapper

  Simple DirectMedia Layer
  Copyright (C) 1997-2013 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/

// #cgo pkg-config: sdl2
// #include <SDL2/SDL.h>
import "C"
import "time"

// ====
// Time
// ====

// Gets the number of milliseconds since the SDL library initialization.
func GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

// Waits a specified number of milliseconds before returning.
func Delay(ms uint32) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func GetPerformanceCounter() uint64 {
	return uint64(C.SDL_GetPerformanceCounter())
}

func GetPerformanceFrequency() uint64 {
	return uint64(C.SDL_GetPerformanceFrequency())
}
