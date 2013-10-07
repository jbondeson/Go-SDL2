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
// #include <SDL2/SDL_image.h>
import "C"
import "unsafe"
import "reflect"

type Surface struct {
	cSurface *C.SDL_Surface

	Flags  uint32
	Format *PixelFormat
	W      int32
	H      int32
	Pitch  int32
	Pixels unsafe.Pointer
	ClipRect Rect

	gcPixels interface{} // Prevents garbage collection of pixels passed to func CreateRGBSurfaceFrom
}

func wrapSurface(cSurface *C.SDL_Surface) *Surface {
	var s *Surface

	if cSurface != nil {
		var surface Surface
		surface.SetCSurface(unsafe.Pointer(cSurface))
		s = &surface
	} else {
		s = nil
	}

	return s
}

// FIXME: Ideally, this should NOT be a public function, but it is needed in the package "ttf" ...
func (s *Surface) SetCSurface(cSurface unsafe.Pointer) {
	s.cSurface = (*C.SDL_Surface)(cSurface)
	s.reload()
}

// Pull data from C.SDL_Surface.
// Make sure to use this when the C surface might have been changed.
func (s *Surface) reload() {
	s.Flags = uint32(s.cSurface.flags)
	s.Format = (*PixelFormat)(cast(s.cSurface.format))
	s.W = int32(s.cSurface.w)
	s.H = int32(s.cSurface.h)
	s.Pitch = int32(s.cSurface.pitch)
	s.Pixels = s.cSurface.pixels
	s.ClipRect.X = int32(s.cSurface.clip_rect.x)
	s.ClipRect.Y = int32(s.cSurface.clip_rect.y)
	s.ClipRect.W = int32(s.cSurface.clip_rect.w)
	s.ClipRect.H = int32(s.cSurface.clip_rect.h)
}

func (s *Surface) destroy() {
	s.cSurface = nil
	s.Format = nil
	s.Pixels = nil
	s.gcPixels = nil
}

// Frees (deletes) a Surface
func (screen *Surface) Free() {
	C.SDL_FreeSurface(screen.cSurface)

	screen.destroy()
}

// Locks a surface for direct access.
func (screen *Surface) Lock() int {
	status := int(C.SDL_LockSurface(screen.cSurface))
	return status
}

// Unlocks a previously locked surface.
func (screen *Surface) Unlock() {
	C.SDL_UnlockSurface(screen.cSurface)
}

// Performs a fast blit from the source surface to the destination surface.
// This is the same as func BlitSurface, but the order of arguments is reversed.
func (dst *Surface) Blit(dstrect *Rect, src *Surface, srcrect *Rect) int {
	var ret C.int
	ret = C.SDL_UpperBlit(
		src.cSurface,
		(*C.SDL_Rect)(cast(srcrect)),
		dst.cSurface,
		(*C.SDL_Rect)(cast(dstrect)))
	return int(ret)
}

// Performs a fast blit from the source surface to the destination surface.
func BlitSurface(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) int {
	return dst.Blit(dstrect, src, srcrect)
}

// This function performs a fast fill of the given rectangle with some color.
func (dst *Surface) FillRect(dstrect *Rect, color uint32) int {
	var ret = C.SDL_FillRect(
		dst.cSurface,
		(*C.SDL_Rect)(cast(dstrect)),
		C.Uint32(color))
	return int(ret)
}

// Sets the color key (transparent pixel)  in  a  blittable  surface  and
// enables or disables RLE blit acceleration.
func (s *Surface) SetColorKey(flags uint32, ColorKey uint32) int {
	status := int(C.SDL_SetColorKey(s.cSurface, C.int(flags), C.Uint32(ColorKey)))
	return status
}

// Gets the clipping rectangle for a surface.
func (s *Surface) GetClipRect(r *Rect) {
	C.SDL_GetClipRect(s.cSurface, (*C.SDL_Rect)(cast(r)))
}

// Sets the clipping rectangle for a surface.
func (s *Surface) SetClipRect(r *Rect) {
	C.SDL_SetClipRect(s.cSurface, (*C.SDL_Rect)(cast(r)))
}

// Loads Surface from file (using IMG_Load).
func Load(file string) *Surface {
	cfile := C.CString(file)
	var screen = C.IMG_Load(cfile)
	C.free(unsafe.Pointer(cfile))
	return wrapSurface(screen)
}

// Creates an empty Surface.
func CreateRGBSurface(flags uint32, width int, height int, bpp int, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) *Surface {
	p := C.SDL_CreateRGBSurface(C.Uint32(flags), C.int(width), C.int(height), C.int(bpp),
		C.Uint32(Rmask), C.Uint32(Gmask), C.Uint32(Bmask), C.Uint32(Amask))
	return wrapSurface(p)
}

// Creates a Surface from existing pixel data. It expects pixels to be a slice, pointer or unsafe.Pointer.
func CreateRGBSurfaceFrom(pixels interface{}, width, height, bpp, pitch int, Rmask, Gmask, Bmask, Amask uint32) *Surface {
	var ptr unsafe.Pointer
	switch v := reflect.ValueOf(pixels); v.Kind() {
	case reflect.Ptr, reflect.UnsafePointer, reflect.Slice:
		ptr = unsafe.Pointer(v.Pointer())
	default:
		panic("Don't know how to handle type: " + v.Kind().String())
	}

	p := C.SDL_CreateRGBSurfaceFrom(ptr, C.int(width), C.int(height), C.int(bpp), C.int(pitch),
		C.Uint32(Rmask), C.Uint32(Gmask), C.Uint32(Bmask), C.Uint32(Amask))

	s := wrapSurface(p)
	s.gcPixels = pixels
	return s
}
