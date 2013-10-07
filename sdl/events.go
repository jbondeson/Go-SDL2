package sdl

/*
  SDL2 Go Wrapper
  Copyright (C) 2013 Kristoffer Gronlund <kgronlund@suse.com>

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

const (
	// event types
	QUIT            = C.SDL_QUIT
	WINDOWEVENT     = C.SDL_WINDOWEVENT
	SYSWMEVENT      = C.SDL_SYSWMEVENT
	KEYDOWN         = C.SDL_KEYDOWN
	KEYUP           = C.SDL_KEYUP
	TEXTEDITING     = C.SDL_TEXTEDITING
	TEXTINPUT       = C.SDL_TEXTINPUT
	MOUSEMOTION     = C.SDL_MOUSEMOTION
	MOUSEBUTTONDOWN = C.SDL_MOUSEBUTTONDOWN
	MOUSEBUTTONUP   = C.SDL_MOUSEBUTTONUP
	MOUSEWHEEL      = C.SDL_MOUSEWHEEL
	JOYAXISMOTION   = C.SDL_JOYAXISMOTION
	JOYBALLMOTION   = C.SDL_JOYBALLMOTION
	JOYHATMOTION    = C.SDL_JOYHATMOTION
	JOYBUTTONDOWN   = C.SDL_JOYBUTTONDOWN
	JOYBUTTONUP     = C.SDL_JOYBUTTONUP
	JOYDEVICEADDED  = C.SDL_JOYDEVICEADDED
	JOYDEVICEREMOVED= C.SDL_JOYDEVICEREMOVED
	CONTROLLERAXISMOTION = C.SDL_CONTROLLERAXISMOTION
	CONTROLLERBUTTONDOWN = C.SDL_CONTROLLERBUTTONDOWN
	CONTROLLERBUTTONUP = C.SDL_CONTROLLERBUTTONUP
	CONTROLLERDEVICEADDED = C.SDL_CONTROLLERDEVICEADDED
	CONTROLLERDEVICEREMOVED = C.SDL_CONTROLLERDEVICEREMOVED
	CONTROLLERDEVICEREMAPPED = C.SDL_CONTROLLERDEVICEREMAPPED
	FINGERDOWN = C.SDL_FINGERDOWN
	FINGERUP = C.SDL_FINGERUP
	FINGERMOTION = C.SDL_FINGERMOTION
	DOLLARGESTURE = C.SDL_DOLLARGESTURE
	DOLLARRECORD = C.SDL_DOLLARRECORD
	MULTIGESTURE = C.SDL_MULTIGESTURE
	CLIPBOARDUPDATE = C.SDL_CLIPBOARDUPDATE
	DROPFILE        = C.SDL_DROPFILE
	USEREVENT       = C.SDL_USEREVENT


	// event state
	QUERY   = C.SDL_QUERY
	IGNORE  = C.SDL_IGNORE
	DISABLE = C.SDL_DISABLE
	ENABLE  = C.SDL_ENABLE

	// constants
	TEXTEDITINGEVENT_TEXT_SIZE = 32
	TEXTINPUTEVENT_TEXT_SIZE = 32
)

// Polls for currently pending events
func (event *Event) Poll() bool {
	ret := C.SDL_PollEvent((*C.SDL_Event)(cast(event)))
	return ret != 0
}

func (event *Event) Pump() {
	C.SDL_PumpEvents()
}

func (event *Event) Wait() bool {
	ret := C.SDL_WaitEvent((*C.SDL_Event)(cast(event)))
	return ret != 0
}

func (event *Event) WaitTimeout(timeout int) bool {
	ret := C.SDL_WaitEventTimeout((*C.SDL_Event)(cast(event)), C.int(timeout))
	return ret != 0
}

// Adapts the event to its type
func (event *Event) Get() interface{} {
	switch event.Type {
	case QUIT:
		return *(*QuitEvent)(cast(event))

	case WINDOWEVENT:
		return *(*WindowEvent)(cast(event))

	case SYSWMEVENT:
		return *(*SysWMEvent)(cast(event))

	case KEYDOWN, KEYUP:
		return *(*KeyboardEvent)(cast(event))

	case TEXTEDITING:
		return *(*TextEditingEvent)(cast(event))

	case TEXTINPUT:
		return *(*TextInputEvent)(cast(event))

	case MOUSEBUTTONDOWN, MOUSEBUTTONUP:
		return *(*MouseButtonEvent)(cast(event))

	case MOUSEMOTION:
		return *(*MouseMotionEvent)(cast(event))

	case MOUSEWHEEL:
		return *(*MouseWheelEvent)(cast(event))

	case JOYAXISMOTION:
		return *(*JoyAxisEvent)(cast(event))

	case JOYBALLMOTION:
		return *(*JoyBallEvent)(cast(event))

	case JOYHATMOTION:
		return *(*JoyHatEvent)(cast(event))

	case JOYBUTTONDOWN, JOYBUTTONUP:
		return *(*JoyButtonEvent)(cast(event))

	case JOYDEVICEADDED, JOYDEVICEREMOVED:
		return *(*JoyDeviceEvent)(cast(event))

	case CLIPBOARDUPDATE:
		return *(*CommonEvent)(cast(event))

	case DROPFILE:
		return *(*DropEvent)(cast(event))

	case USEREVENT:
		return *(*UserEvent)(cast(event))

	}

	return nil
}
