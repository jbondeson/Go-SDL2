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


// #cgo pkg-config: sdl2 SDL2_image
//
// struct private_hwdata{};
// //struct SDL_BlitMap{};
// #define map _map
//
// #include <SDL2/SDL.h>
import "C"

import (
	"os"
	"runtime"
	"reflect"
	"unsafe"
)

type cast unsafe.Pointer

func ptr(v interface{}) unsafe.Pointer {

	if v == nil {
		return unsafe.Pointer(nil)
	}

	rv := reflect.ValueOf(v)
	var et reflect.Value
	switch rv.Type().Kind() {
	case reflect.Uintptr:
		offset, _ := v.(uintptr)
		return unsafe.Pointer(offset)
	case reflect.Ptr:
		et = rv.Elem()
	case reflect.Slice:
		et = rv.Index(0)
	default:
		panic("type must be a pointer, a slice, uintptr or nil")
	}

	return unsafe.Pointer(et.UnsafeAddr())
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

// ======
// SDL.h
// =====

// The version of Go-SDL bindings.
// The version descriptor changes into a new unique string
// after a semantically incompatible Go-SDL update.
//
// The returned value can be checked by users of this package
// to make sure they are using a version with the expected semantics.
//
// If Go adds some kind of support for package versioning, this function will go away.
func GoSdlVersion() string {
	return "krig SDL bindings 1.0"
}

// Initializes SDL.
func Init(flags uint32) int {
	status := int(C.SDL_Init(C.Uint32(flags)))
	if (status != 0) && (runtime.GOOS == "darwin") && (flags&INIT_VIDEO != 0) {
		if os.Getenv("SDL_VIDEODRIVER") == "" {
			os.Setenv("SDL_VIDEODRIVER", "x11")
			status = int(C.SDL_Init(C.Uint32(flags)))
			if status != 0 {
				os.Setenv("SDL_VIDEODRIVER", "")
			}
		}
	}
	return status
}

// Shuts down SDL
func Quit() {
	C.SDL_Quit()
}

// Initializes subsystems.
func InitSubSystem(flags uint32) int {
	status := int(C.SDL_InitSubSystem(C.Uint32(flags)))
	if (status != 0) && (runtime.GOOS == "darwin") && (flags&INIT_VIDEO != 0) {
		if os.Getenv("SDL_VIDEODRIVER") == "" {
			os.Setenv("SDL_VIDEODRIVER", "x11")
			status = int(C.SDL_InitSubSystem(C.Uint32(flags)))
			if status != 0 {
				os.Setenv("SDL_VIDEODRIVER", "")
			}
		}
	}
	return status
}

// Shuts down a subsystem.
func QuitSubSystem(flags uint32) {
	C.SDL_QuitSubSystem(C.Uint32(flags))
}

// Checks which subsystems are initialized.
func WasInit(flags uint32) uint32 {
	status := uint32(C.SDL_WasInit(C.Uint32(flags)))
	return status
}
