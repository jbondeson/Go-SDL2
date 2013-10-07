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
import "unsafe"

// =========
// Keyboard
// ======

// Gets the state of modifier keys
func GetModState() int32 {
	state := int32(C.SDL_GetModState())
	return state
}

// Sets the state of modifier keys
func SetModState(modstate int32) {
	C.SDL_SetModState(C.SDL_Keymod(modstate))
}

// Gets the name of an SDL virtual keysym
func GetKeyName(key int32) string {
	name := C.GoString(C.SDL_GetKeyName(C.SDL_Keycode(key)))
	return name
}


// =====
// Mouse
// =====

func GetMouseFocus() *Window {
	window := C.SDL_GetMouseFocus()
	return wrapWindow(window)
}

// Retrieves the current state of the mouse.
func GetMouseState(x, y *int) uint32 {
	state := uint32(C.SDL_GetMouseState((*C.int)(cast(x)), (*C.int)(cast(y))))
	return state
}

// Retrieves the current state of the mouse relative to the last time this
// function was called.
func GetRelativeMouseState(x, y *int) uint32 {
	state := uint32(C.SDL_GetRelativeMouseState((*C.int)(cast(x)), (*C.int)(cast(y))))
	return state
}

/**
 *  \brief Moves the mouse to the given position within the window.
 *
 *  \param window The window to move the mouse into, or NULL for the current mouse focus
 *  \param x The x coordinate within the window
 *  \param y The y coordinate within the window
 *
 *  \note This function generates a mouse motion event
 */
func WarpMouseInWindow(window *Window, x int, y int) {
	C.SDL_WarpMouseInWindow(window.cWindow, C.int(x), C.int(y))
}


// Toggle whether or not the cursor is shown on the screen.
func ShowCursor(toggle int) int {
	state := int(C.SDL_ShowCursor((C.int)(toggle)))
	return state
}

// ========
// Joystick
// ========

type Joystick struct {
	cJoystick *C.SDL_Joystick
}

func wrapJoystick(cJoystick *C.SDL_Joystick) *Joystick {
	var j *Joystick
	if cJoystick != nil {
		var joystick Joystick
		joystick.cJoystick = (*C.SDL_Joystick)(unsafe.Pointer(cJoystick))
		j = &joystick
	} else {
		j = nil
	}
	return j
}

// Count the number of joysticks attached to the system
func NumJoysticks() int {
	num := int(C.SDL_NumJoysticks())
	return num
}

// Open a joystick for use The index passed as an argument refers to
// the N'th joystick on the system. This index is the value which will
// identify this joystick in future joystick events.  This function
// returns a joystick identifier, or NULL if an error occurred.
func JoystickOpen(deviceIndex int) *Joystick {
	joystick := C.SDL_JoystickOpen(C.int(deviceIndex))
	return wrapJoystick(joystick)
}

// Update the current state of the open joysticks. This is called
// automatically by the event loop if any joystick events are enabled.
func JoystickUpdate() {
	C.SDL_JoystickUpdate()
}

// Enable/disable joystick event polling. If joystick events are
// disabled, you must call SDL_JoystickUpdate() yourself and check the
// state of the joystick when you want joystick information. The state
// can be one of SDL_QUERY, SDL_ENABLE or SDL_IGNORE.
func JoystickEventState(state int) int {
	result := int(C.SDL_JoystickEventState(C.int(state)))
	return result
}

// Close a joystick previously opened with SDL_JoystickOpen()
func (joystick *Joystick) Close() {
	C.SDL_JoystickClose(joystick.cJoystick)
}

// Get the number of general axis controls on a joystick
func (joystick *Joystick) NumAxes() int {
	return int(C.SDL_JoystickNumAxes(joystick.cJoystick))
}

// Get the number of buttons on a joystick
func (joystick *Joystick) NumButtons() int {
	return int(C.SDL_JoystickNumButtons(joystick.cJoystick))
}

// Get the number of trackballs on a Joystick trackballs have only
// relative motion events associated with them and their state cannot
// be polled.
func (joystick *Joystick) NumBalls() int {
	return int(C.SDL_JoystickNumBalls(joystick.cJoystick))
}

// Get the number of POV hats on a joystick
func (joystick *Joystick) NumHats() int {
	return int(C.SDL_JoystickNumHats(joystick.cJoystick))
}

// Get the current state of a POV hat on a joystick
// The hat indices start at index 0.
func (joystick *Joystick) GetHat(hat int) uint8 {
	return uint8(C.SDL_JoystickGetHat(joystick.cJoystick, C.int(hat)))
}

// Get the current state of a button on a joystick. The button indices
// start at index 0.
func (joystick *Joystick) GetButton(button int) uint8 {
	return uint8(C.SDL_JoystickGetButton(joystick.cJoystick, C.int(button)))
}

// Get the ball axis change since the last poll. The ball indices
// start at index 0. This returns 0, or -1 if you passed it invalid
// parameters.
func (joystick *Joystick) GetBall(ball int, dx, dy *int) int {
	return int(C.SDL_JoystickGetBall(joystick.cJoystick, C.int(ball), (*C.int)(cast(dx)), (*C.int)(cast(dy))))
}

// Get the current state of an axis control on a joystick. The axis
// indices start at index 0. The state is a value ranging from -32768
// to 32767.
func (joystick *Joystick) GetAxis(axis int) int16 {
	return int16(C.SDL_JoystickGetAxis(joystick.cJoystick, C.int(axis)))
}
