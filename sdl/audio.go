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

const (
	AUDIO_MASK_BITSIZE = C.SDL_AUDIO_MASK_BITSIZE
	AUDIO_MASK_DATATYPE = C.SDL_AUDIO_MASK_DATATYPE
	AUDIO_MASK_ENDIAN = C.SDL_AUDIO_MASK_ENDIAN
	AUDIO_MASK_SIGNED = C.SDL_AUDIO_MASK_SIGNED
	AUDIO_U8 = C.AUDIO_U8
	AUDIO_S8 = C.AUDIO_S8
	AUDIO_U16LSB = C.AUDIO_U16LSB
	AUDIO_S16LSB = C.AUDIO_S16LSB
	AUDIO_U16MSB = C.AUDIO_U16MSB
	AUDIO_S16MSB = C.AUDIO_S16MSB
	AUDIO_U16 = C.AUDIO_U16
	AUDIO_S16 = C.AUDIO_S16
	AUDIO_S32LSB = C.AUDIO_S32LSB
	AUDIO_S32MSB = C.AUDIO_S32MSB
	AUDIO_S32 = C.AUDIO_S32
	AUDIO_F32LSB = C.AUDIO_F32LSB
	AUDIO_F32MSB = C.AUDIO_F32MSB
	AUDIO_F32 = C.AUDIO_F32
	AUDIO_U16SYS = C.AUDIO_U16SYS
	AUDIO_S16SYS = C.AUDIO_S16SYS
	AUDIO_S32SYS = C.AUDIO_S32SYS
	AUDIO_F32SYS = C.AUDIO_F32SYS

	AUDIO_ALLOW_FREQUENCY_CHANGE = C.SDL_AUDIO_ALLOW_FREQUENCY_CHANGE
	AUDIO_ALLOW_FORMAT_CHANGE = C.SDL_AUDIO_ALLOW_FORMAT_CHANGE
	AUDIO_ALLOW_CHANNELS_CHANGE = C.SDL_AUDIO_ALLOW_CHANNELS_CHANGE
	AUDIO_ALLOW_ANY_CHANGE = C.SDL_AUDIO_ALLOW_ANY_CHANGE
)

func Audio_BitSize(x uint32) uint32 {
	return x & AUDIO_MASK_BITSIZE
}

func Audio_IsFloat(x uint32) bool {
	return (x & AUDIO_MASK_DATATYPE) != 0
}

func Audio_IsBigEndian(x uint32) bool {
	return (x & AUDIO_MASK_ENDIAN) != 0
}

func Audio_IsSigned(x uint32) bool {
	return (x & AUDIO_MASK_SIGNED) != 0
}

func Audio_IsInt(x uint32) bool {
	return !Audio_IsFloat(x)
}

func Audio_IsLittleEndian(x uint32) bool {
	return !Audio_IsBigEndian(x)
}

func Audio_IsUnsigned(x uint32) bool {
	return !Audio_IsSigned(x)
}


type AudioCallback func(userdata unsafe.Pointer, stream *byte, len int32)

type AudioSpec struct {
	Freq int32
	Format uint16
	Channels uint8
	Silence uint8
	Samples uint16
	Padding uint16
	Size uint32
	Callback AudioCallback
	UserData interface{}
}

func GetNumAudioDrivers() int {
	return int(C.SDL_GetNumAudioDrivers())
}

func GetAudioDriver(index int) string {
	return C.GoString(C.SDL_GetAudioDriver(C.int(index)))
}

func GetCurrentAudioDriver() string {
	return C.GoString(C.SDL_GetCurrentAudioDriver())
}

//func OpenAudio(desired AudioSpec) (AudioSpec, int) {
//	// magic, wrapAudioSpec
//	ret := int(C.SDL_OpenAudio(c_desired, c_obtained))
//	// more magic, unwrapAudioSpec
//	return obtained, ret
//}

func GetNumAudioDevices(iscapture int) int {
	return int(C.SDL_GetNumAudioDevices(C.int(iscapture)))
}

func GetAudioDeviceName(index, iscapture int) string {
	return C.GoString(C.SDL_GetAudioDeviceName(C.int(index), C.int(iscapture)))
}

func GetAudioStatus() int32 {
	return int32(C.SDL_GetAudioStatus())
}

func GetAudioDeviceStatus(dev uint32) int32 {
	return int32(C.SDL_GetAudioDeviceStatus(C.SDL_AudioDeviceID(dev)))
}

func PauseAudio(pause_on int) {
	C.SDL_PauseAudio(C.int(pause_on))
}

func PauseAudioDevice(dev uint32, pause_on int) {
	C.SDL_PauseAudioDevice(C.SDL_AudioDeviceID(dev), C.int(pause_on))
}

type WAVData struct {
	Spec *AudioSpec
	AudioBuf *byte
	AudioLen uint32
}

func LoadWAV(file string, spec *AudioSpec) *WAVData {
	// magic
	return nil
}

func (w *WAVData) Free() {
	C.SDL_FreeWAV((*C.Uint8)(unsafe.Pointer(w.AudioBuf)))
}

func LockAudio() {
	C.SDL_LockAudio()
}

func UnlockAudio() {
	C.SDL_UnlockAudio()
}

func LockAudioDevice(dev uint32) {
	C.SDL_LockAudioDevice(C.SDL_AudioDeviceID(dev))
}

func UnlockAudioDevice(dev uint32) {
	C.SDL_UnlockAudioDevice(C.SDL_AudioDeviceID(dev))
}