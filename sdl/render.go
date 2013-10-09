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

type Renderer struct {
	cRenderer *C.SDL_Renderer
}

type Texture struct {
	cTexture *C.SDL_Texture
}

func wrapRenderer(cRenderer *C.SDL_Renderer) *Renderer {
	var r *Renderer

	if cRenderer != nil {
		var renderer Renderer
		renderer.cRenderer = (*C.SDL_Renderer)(unsafe.Pointer(cRenderer))
		r = &renderer
	} else {
		r = nil
	}

	return r
}

func wrapTexture(cTexture *C.SDL_Texture) *Texture {
	var t *Texture

	if cTexture != nil {
		var texture Texture
		texture.cTexture = (*C.SDL_Texture)(unsafe.Pointer(cTexture))
		t = &texture
	} else {
		t = nil
	}

	return t
}

func CreateRenderer(w *Window, index int, flags uint32) *Renderer {
	renderer := C.SDL_CreateRenderer(w.cWindow, C.int(index), C.Uint32(flags))

	return wrapRenderer(renderer)
}

func (r *Renderer) Clear() {
	C.SDL_RenderClear(r.cRenderer)
}

func (r *Renderer) Present() {
	C.SDL_RenderPresent(r.cRenderer)
}

func (r *Renderer) GetDrawColor() Color {
	cr := C.Uint8(0)
	cg := C.Uint8(0)
	cb := C.Uint8(0)
	ca := C.Uint8(0)
	ret := C.SDL_GetRenderDrawColor(r.cRenderer, &cr, &cg, &cb, &ca)
	if int(ret) == 0 {
		return Color{uint8(cr), uint8(cg), uint8(cb), uint8(ca)}
	}
	return Color{0, 0, 0, 0}
}

func (r *Renderer) SetDrawColor(c Color) {
	C.SDL_SetRenderDrawColor(r.cRenderer, C.Uint8(c.R),
		C.Uint8(c.G), C.Uint8(c.B), C.Uint8(c.A))
}

func (r *Renderer) Destroy() {
	C.SDL_DestroyRenderer(r.cRenderer)
}

func (r *Renderer) DrawPoint(x, y int) bool {
	ret := C.SDL_RenderDrawPoint(r.cRenderer, C.int(x), C.int(y))
	return int(ret) == 0
}

func (r *Renderer) DrawPoints(points []Point) bool {
	pptr := unsafe.Pointer(&points[0])
	n := len(points)
	ret := C.SDL_RenderDrawPoints(r.cRenderer, (*C.SDL_Point)(pptr), C.int(n))
	return int(ret) == 0
}

func (r *Renderer) DrawLine(x1, y1, x2, y2 int) bool {
	ret := C.SDL_RenderDrawLine(r.cRenderer, C.int(x1), C.int(y1), C.int(x2), C.int(y2))
	return int(ret) == 0
}

func (r *Renderer) DrawLines(points []Point) bool {
	pptr := unsafe.Pointer(&points[0])
	n := len(points)
	ret := C.SDL_RenderDrawLines(r.cRenderer, (*C.SDL_Point)(pptr), C.int(n))
	return int(ret) == 0
}

func (r *Renderer) DrawRect(rect *Rect) bool {
	ret := C.SDL_RenderDrawRect(r.cRenderer, (*C.SDL_Rect)(cast(rect)))
	return int(ret) == 0
}

func (r *Renderer) DrawRects(rects []Rect) bool {
	prects := unsafe.Pointer(&rects[0])
	n := len(rects)
	ret := C.SDL_RenderDrawRects(r.cRenderer, (*C.SDL_Rect)(prects), C.int(n))
	return int(ret) == 0
}

func (r *Renderer) FillRect(rect *Rect) bool {
	ret := C.SDL_RenderFillRect(r.cRenderer, (*C.SDL_Rect)(cast(rect)))
	return int(ret) == 0
}

func (r *Renderer) FillRects(rects []Rect) bool {
	prects := unsafe.Pointer(&rects[0])
	n := len(rects)
	ret := C.SDL_RenderFillRects(r.cRenderer, (*C.SDL_Rect)(prects), C.int(n))
	return int(ret) == 0
}

func (r *Renderer) SetDrawBlendMode(blendmode int) bool {
	ret := C.SDL_SetRenderDrawBlendMode(r.cRenderer, C.SDL_BlendMode(blendmode))
	return int(ret) == 0
}

func (r *Renderer) GetDrawBlendMode() (int, bool) {
	bm := C.SDL_BlendMode(0)
	ret := C.SDL_GetRenderDrawBlendMode(r.cRenderer, &bm)
	return int(bm), int(ret) == 0
}

func (r *Renderer) Copy(t *Texture, src, dst *Rect) bool {
	ret := C.SDL_RenderCopy(r.cRenderer, t.cTexture,
		(*C.SDL_Rect)(cast(src)), (*C.SDL_Rect)(cast(dst)))
	return int(ret) == 0
}

func (r *Renderer) CopyEx(t *Texture, src, dst *Rect, angle float64, center *Point, flip int) bool {
	ret := C.SDL_RenderCopyEx(r.cRenderer, t.cTexture,
		(*C.SDL_Rect)(cast(src)), (*C.SDL_Rect)(cast(dst)),
		C.double(angle), (*C.SDL_Point)(cast(center)), C.SDL_RendererFlip(flip))
	return int(ret) == 0
}


func CreateWindowAndRenderer(h, w int, flags uint32) (*Window, *Renderer) {
	var win Window
	var rend Renderer

	C.SDL_CreateWindowAndRenderer(C.int(h), C.int(w), C.Uint32(flags),
		&win.cWindow, &rend.cRenderer)

	return &win, &rend
}

func (r *Renderer) CreateTexture(format uint32, access, w, h int) *Texture {
	texture := C.SDL_CreateTexture(r.cRenderer, C.Uint32(format),
		C.int(access), C.int(w), C.int(h))

	return wrapTexture(texture)
}

func (r *Renderer) CreateTextureFromSurface(s *Surface) *Texture {
	texture := C.SDL_CreateTextureFromSurface(r.cRenderer, s.cSurface)
	return wrapTexture(texture)
}

func (t *Texture) Update(rect *Rect, pixels interface{}, pitch int) {
	C.SDL_UpdateTexture(t.cTexture, (*C.SDL_Rect)(cast(rect)), ptr(pixels), C.int(pitch))
}

func (t *Texture) Destroy() {
	C.SDL_DestroyTexture(t.cTexture)
}

// Returns (ok, texture width, texture height)
func (t *Texture) Bind() (float32, float32, bool) {
	texw := C.float(0.0)
	texh := C.float(0.0)
	ret := C.SDL_GL_BindTexture(t.cTexture, &texw, &texh)
	return float32(texw), float32(texh), int(ret) == 0
}

func (t *Texture) Unbind() bool {
	return int(C.SDL_GL_UnbindTexture(t.cTexture)) == 0
}

// Query returns (w, h)
func (t *Texture) GetSize() (int, int) {
	var w C.int
	var h C.int
	C.SDL_QueryTexture(t.cTexture, nil, nil, &w, &h)
	return int(w), int(h)
}

func (r *Rect) Empty() bool {
	return (r.W <= 0) || (r.H <= 0)
}

func (a *Rect) Equals(b *Rect) bool {
	return (a.X == b.X) && (a.Y == b.Y) && (a.W == b.W) && (a.H == b.H)
}

func (a *Rect) HasIntersection(b *Rect) bool {
	return C.SDL_HasIntersection((*C.SDL_Rect)(unsafe.Pointer(a)), (*C.SDL_Rect)(unsafe.Pointer(b))) == C.SDL_TRUE
}

func (r *Rect) Contains(x, y int32) bool {
	return !(x < r.X || x > (r.X + r.W) || y < r.Y || y > (r.Y + r.H))
}

func (a *Rect) Intersect(b *Rect) *Rect {
	var ret Rect
	is := C.SDL_IntersectRect((*C.SDL_Rect)(unsafe.Pointer(a)), (*C.SDL_Rect)(unsafe.Pointer(b)), (*C.SDL_Rect)(unsafe.Pointer(&ret)))
	if is == C.SDL_TRUE {
		return &ret
	}
	return nil
}

func (a *Rect) Union(b *Rect) *Rect {
	var ret Rect
	C.SDL_UnionRect((*C.SDL_Rect)(unsafe.Pointer(a)), (*C.SDL_Rect)(unsafe.Pointer(b)), (*C.SDL_Rect)(unsafe.Pointer(&ret)))
	return &ret
}

func (clip *Rect) Enclose(points []Point) *Rect {
	var ret Rect
	is := C.SDL_EnclosePoints((*C.SDL_Point)(unsafe.Pointer(&points[0])), C.int(len(points)), (*C.SDL_Rect)(unsafe.Pointer(clip)), (*C.SDL_Rect)(unsafe.Pointer(&ret)))
	if is == C.SDL_TRUE {
		return &ret
	}
	return nil
}

// Calculate the intersection of a rectangle and line segment.
// Return SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
func (r *Rect) IntersectLine(x1, y1, x2, y2 *int32) bool  {
	ret := C.SDL_IntersectRectAndLine((*C.SDL_Rect)(unsafe.Pointer(r)), (*C.int)(unsafe.Pointer(x1)), (*C.int)(unsafe.Pointer(y1)), (*C.int)(unsafe.Pointer(x2)), (*C.int)(unsafe.Pointer(y2)))
	return ret == C.SDL_TRUE
}
