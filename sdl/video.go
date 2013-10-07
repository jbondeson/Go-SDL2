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

type Window struct {
	cWindow *C.SDL_Window
}

func wrapWindow(cWindow *C.SDL_Window) *Window {
	var w *Window

	if cWindow != nil {
		var window Window
		w.cWindow = (*C.SDL_Window)(cWindow)
		w = &window
	} else {
		w = nil
	}

	return w
}

func CreateWindow(title string, x, y int, w, h int, flags uint32) *Window {
	window := C.SDL_CreateWindow(C.CString(title), C.int(x), C.int(y),
		C.int(w), C.int(h), C.Uint32(flags))

	return wrapWindow(window)
}

// Swaps OpenGL framebuffers/Update Display.
func (w *Window) GL_SwapWindow() {
	C.SDL_GL_SwapWindow(w.cWindow)
}

func (w *Window) GL_CreateContext() {
	C.SDL_GL_CreateContext(w.cWindow)
}

func GL_SetAttribute(attr int, value int) int {
	status := int(C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), C.int(value)))
	return status
}

func NumDisplayModes(index int) int {
	return int(C.SDL_GetNumDisplayModes(C.int(index)))
}

func (w *Window) GetTitle() string {
	ctitle := C.SDL_GetWindowTitle(w.cWindow)

	return C.GoString(ctitle)
}

func (w *Window) SetTitle(title string) {
	ctitle := C.CString(title)
	C.SDL_SetWindowTitle(w.cWindow, ctitle)

	C.free(unsafe.Pointer(ctitle))
}

func (w *Window) SetIcon(s *Surface) {
	C.SDL_SetWindowIcon(w.cWindow, s.cSurface)
}

func (w *Window) SetFullscreen(flags uint32) {
	C.SDL_SetWindowFullscreen(w.cWindow, C.Uint32(flags))
}

func (w *Window) GetSurface() *Surface {
	return wrapSurface(C.SDL_GetWindowSurface(w.cWindow))
}

func (w *Window) UpdateSurface() int {
	return int(C.SDL_UpdateWindowSurface(w.cWindow))
}

func (w *Window) Destroy() {
	C.SDL_DestroyWindow(w.cWindow)
}

func (w *Window) GetSize() (int, int) {
	cw := C.int(0)
	ch := C.int(0)
	C.SDL_GetWindowSize(w.cWindow, &cw, &ch)
	return int(cw), int(ch)
}

func (w *Window) GetMinimumSize() (int, int) {
	cw := C.int(0)
	ch := C.int(0)
	C.SDL_GetWindowMinimumSize(w.cWindow, &cw, &ch)
	return int(cw), int(ch)
}

func (w *Window) ShowSimpleMessageBox(flags uint32, title, message string) {
	ctitle, cmessage := C.CString(title), C.CString(message)
	C.SDL_ShowSimpleMessageBox(C.Uint32(flags), ctitle, cmessage, w.cWindow)

	C.free(unsafe.Pointer(ctitle))
	C.free(unsafe.Pointer(cmessage))
}

func (w *Window) SetGrab(grabbed bool) {
	C.SDL_SetWindowGrab(w.cWindow, C.SDL_bool(bool2int(grabbed)))
}

func EnableScreenSaver() {
	C.SDL_EnableScreenSaver()
}

func DisableScreenSaver() {
	C.SDL_DisableScreenSaver()
}

func IsScreenSaverEnabled() bool {
	return int(C.SDL_IsScreenSaverEnabled()) != 0
}

// Map a RGBA color value to a pixel format.
func MapRGBA(format *PixelFormat, r, g, b, a uint8) uint32 {
	return (uint32)(C.SDL_MapRGBA((*C.SDL_PixelFormat)(cast(format)), (C.Uint8)(r), (C.Uint8)(g), (C.Uint8)(b), (C.Uint8)(a)))
}

// Gets RGBA values from a pixel in the specified pixel format.
func GetRGBA(color uint32, format *PixelFormat, r, g, b, a *uint8) {
	C.SDL_GetRGBA(C.Uint32(color), (*C.SDL_PixelFormat)(cast(format)), (*C.Uint8)(r), (*C.Uint8)(g), (*C.Uint8)(b), (*C.Uint8)(a))
}
