package sdl

/*
  SDL2 Go Wrapper
  Copyright (C) 2013 Kristoffer Gronlund <kgronlund@suse.com>
  Copyright (C) Scott Ferguson <user scottferg at github>
  Copyright (C) Piotr Praszmo <user banthar at github>

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
//
// #include <SDL2/SDL.h>
import "C"
import "unsafe"

type RWops struct {
	cRWops *C.SDL_RWops
	mem    []byte // Retain reference to memory passed to RWFromMem
}

func wrapRWops(cRWops *C.SDL_RWops) *RWops {
	var r *RWops

	if cRWops != nil {
		var ops RWops
		ops.cRWops = (*C.SDL_RWops)(unsafe.Pointer(cRWops))
		ops.mem = nil
		r = &ops
	} else {
		r = nil
	}
	return r
}

func (rwops *RWops) Free() {
	C.SDL_FreeRW(rwops.cRWops)
	rwops.cRWops = nil
	rwops.mem = nil
}

func RWFromFile(file string, mode string) *RWops {
	cfile, cmode := C.CString(file), C.CString(mode)
	defer C.free(unsafe.Pointer(cfile))
	defer C.free(unsafe.Pointer(cmode))
	return wrapRWops(C.SDL_RWFromFile(cfile, cmode))
}

func RWFromMem(mem []byte) *RWops {
	rw := wrapRWops(C.SDL_RWFromMem(unsafe.Pointer(&mem[0]), C.int(len(mem))))
	rw.mem = mem
	return rw
}
