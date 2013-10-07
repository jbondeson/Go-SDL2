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

func GetCPUCount() int {
	return int(C.SDL_GetCPUCount())
}

func GetCPUCacheLineSize() int {
	return int(C.SDL_GetCPUCacheLineSize())
}

/**
 *  This function returns true if the CPU has AltiVec features.
 */
func HasAltiVec() bool {
	return int(C.SDL_HasAltiVec()) != 0
}

/**
 *  This function returns true if the CPU has MMX features.
 */
func HasMMX() bool {
	return int(C.SDL_HasMMX()) != 0
}

/**
 *  This function returns true if the CPU has 3DNow! features.
 */
func Has3DNow() bool {
	return int(C.SDL_Has3DNow()) != 0
}

/**
 *  This function returns true if the CPU has SSE features.
 */
func HasSSE() bool {
	return int(C.SDL_HasSSE()) != 0
}

/**
 *  This function returns true if the CPU has SSE2 features.
 */
func HasSSE2() bool {
	return int(C.SDL_HasSSE2()) != 0
}

/**
 *  This function returns true if the CPU has SSE3 features.
 */
func HasSSE3() bool {
	return int(C.SDL_HasSSE3()) != 0
}

/**
 *  This function returns true if the CPU has SSE4.1 features.
 */
func HasSSE41() bool {
	return int(C.SDL_HasSSE41()) != 0
}

/**
 *  This function returns true if the CPU has SSE4.2 features.
 */
func HasSSE42() bool {
	return int(C.SDL_HasSSE42()) != 0
}
