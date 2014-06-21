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

import "unsafe"

/*
  #cgo pkg-config: sdl2
  #cgo linux LDFLAGS: -lrt
  #include <SDL2/SDL_audio.h>

  extern void go_sdl2_audio_callback(void* userdata, Uint8* stream, int len);

  SDL_AudioCallback go_sdl2_get_callback();
*/
import "C"

//export go_sdl2_audio_callback
func go_sdl2_audio_callback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	cb := *(*func(*byte, int))(userdata)
	cb((*byte)(stream), int(length))
}

type AudioDeviceID uint32
type AudioFormat uint16

const (
	AUDIO_MASK_BITSIZE  = C.SDL_AUDIO_MASK_BITSIZE
	AUDIO_MASK_DATATYPE = C.SDL_AUDIO_MASK_DATATYPE
	AUDIO_MASK_ENDIAN   = C.SDL_AUDIO_MASK_ENDIAN
	AUDIO_MASK_SIGNED   = C.SDL_AUDIO_MASK_SIGNED
	AUDIO_U8            = AudioFormat(C.AUDIO_U8)
	AUDIO_S8            = AudioFormat(C.AUDIO_S8)
	AUDIO_U16LSB        = AudioFormat(C.AUDIO_U16LSB)
	AUDIO_S16LSB        = AudioFormat(C.AUDIO_S16LSB)
	AUDIO_U16MSB        = AudioFormat(C.AUDIO_U16MSB)
	AUDIO_S16MSB        = AudioFormat(C.AUDIO_S16MSB)
	AUDIO_U16           = AudioFormat(C.AUDIO_U16)
	AUDIO_S16           = AudioFormat(C.AUDIO_S16)
	AUDIO_S32LSB        = AudioFormat(C.AUDIO_S32LSB)
	AUDIO_S32MSB        = AudioFormat(C.AUDIO_S32MSB)
	AUDIO_S32           = AudioFormat(C.AUDIO_S32)
	AUDIO_F32LSB        = AudioFormat(C.AUDIO_F32LSB)
	AUDIO_F32MSB        = AudioFormat(C.AUDIO_F32MSB)
	AUDIO_F32           = AudioFormat(C.AUDIO_F32)
	AUDIO_U16SYS        = AudioFormat(C.AUDIO_U16SYS)
	AUDIO_S16SYS        = AudioFormat(C.AUDIO_S16SYS)
	AUDIO_S32SYS        = AudioFormat(C.AUDIO_S32SYS)
	AUDIO_F32SYS        = AudioFormat(C.AUDIO_F32SYS)

	AUDIO_ALLOW_FREQUENCY_CHANGE = C.SDL_AUDIO_ALLOW_FREQUENCY_CHANGE
	AUDIO_ALLOW_FORMAT_CHANGE    = C.SDL_AUDIO_ALLOW_FORMAT_CHANGE
	AUDIO_ALLOW_CHANNELS_CHANGE  = C.SDL_AUDIO_ALLOW_CHANNELS_CHANGE
	AUDIO_ALLOW_ANY_CHANGE       = C.SDL_AUDIO_ALLOW_ANY_CHANGE

	MIX_MAXVOLUME = 128
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

type AudioSpec struct {
	Freq     int32
	Format   AudioFormat
	Channels uint8
	Silence  uint8
	Samples  uint16
	Padding  uint16
	Size     uint32
}

func AudioInit(driver string) bool {
	cdriver := C.CString(driver)
	ret := C.SDL_AudioInit(cdriver)
	C.free(unsafe.Pointer(cdriver))
	return int(ret) != 0
}

func AudioQuit() {
	C.SDL_AudioQuit()
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

func OpenAudio(desired, obtained *AudioSpec, callback func(*byte, int)) bool {
	var cdesired C.SDL_AudioSpec
	var cobtained C.SDL_AudioSpec
	cdesired.freq = C.int(desired.Freq)
	cdesired.format = C.SDL_AudioFormat(int(desired.Format))
	cdesired.channels = C.Uint8(desired.Channels)
	cdesired.samples = C.Uint16(desired.Samples)
	cdesired.size = C.Uint32(0)
	cdesired.callback = C.go_sdl2_get_callback()
	cdesired.userdata = unsafe.Pointer(&callback)
	ret := C.SDL_OpenAudio(&cdesired, &cobtained)
	if obtained != nil {
		obtained.Freq = int32(cobtained.freq)
		obtained.Format = AudioFormat(int(cobtained.format))
		obtained.Channels = uint8(cobtained.channels)
		obtained.Silence = uint8(cobtained.silence)
		obtained.Samples = uint16(cobtained.samples)
		obtained.Size = uint32(cobtained.size)
	}
	return int(ret) == 0
}

func OpenAudioDevice(device string, iscapture bool, desired, obtained *AudioSpec, allowed_changes bool, callback func(*byte, int)) bool {
	var cdesired C.SDL_AudioSpec
	var cobtained C.SDL_AudioSpec
	cdevice := C.CString(device)
	defer C.free(unsafe.Pointer(cdevice))
	cdesired.freq = C.int(desired.Freq)
	cdesired.format = C.SDL_AudioFormat(int(desired.Format))
	cdesired.channels = C.Uint8(desired.Channels)
	cdesired.samples = C.Uint16(desired.Samples)
	cdesired.size = C.Uint32(0)
	cdesired.callback = C.go_sdl2_get_callback()
	cdesired.userdata = unsafe.Pointer(&callback)
	ret := C.SDL_OpenAudioDevice(cdevice, C.int(bool2int(iscapture)), &cdesired, &cobtained, C.int(bool2int(allowed_changes)))
	if obtained != nil {
		obtained.Freq = int32(cobtained.freq)
		obtained.Format = AudioFormat(int(cobtained.format))
		obtained.Channels = uint8(cobtained.channels)
		obtained.Silence = uint8(cobtained.silence)
		obtained.Samples = uint16(cobtained.samples)
		obtained.Size = uint32(cobtained.size)
	}
	return int(ret) == 0
}

func GetNumAudioDevices(iscapture int) int {
	return int(C.SDL_GetNumAudioDevices(C.int(iscapture)))
}

func GetAudioDeviceName(index int, iscapture bool) string {
	return C.GoString(C.SDL_GetAudioDeviceName(C.int(index), C.int(bool2int(iscapture))))
}

func GetAudioStatus() int {
	return int(C.SDL_GetAudioStatus())
}

func GetAudioDeviceStatus(dev AudioDeviceID) int {
	return int(C.SDL_GetAudioDeviceStatus(C.SDL_AudioDeviceID(int(dev))))
}

func PauseAudio(pause_on bool) {
	C.SDL_PauseAudio(C.int(bool2int(pause_on)))
}

func PauseAudioDevice(dev AudioDeviceID, pause_on bool) {
	C.SDL_PauseAudioDevice(C.SDL_AudioDeviceID(int(dev)), C.int(bool2int(pause_on)))
}

type WAVData struct {
	Spec     *AudioSpec
	AudioBuf *byte
	AudioLen uint32
}

func LoadWAV(file string, spec *AudioSpec) *WAVData {
	// magic
	return nil
}

func LoadWAV_RW(src *RWops, freesrc bool, spec *AudioSpec) *WAVData {
	return nil
}

func (w *WAVData) Free() {
	C.SDL_FreeWAV((*C.Uint8)(unsafe.Pointer(w.AudioBuf)))
}

type AudioFilter func(*AudioCVT, AudioFormat)

type AudioCVT struct {
	cAudioCVT *C.SDL_AudioCVT
	Buf       []byte
}

// returns (cvt, 0/1/-1) -- 0 means no conversion, 1 means filter is set up
func BuildAudioCVT(src_format AudioFormat, src_channels uint8, src_rate int,
	dst_format AudioFormat, dst_channels uint8, dst_rate int) (*AudioCVT, int) {
	var cvt C.SDL_AudioCVT
	ret := C.SDL_BuildAudioCVT(&cvt,
		C.SDL_AudioFormat(int(src_format)), C.Uint8(src_channels), C.int(src_rate),
		C.SDL_AudioFormat(int(dst_format)), C.Uint8(dst_channels), C.int(dst_rate))
	rcvt := &AudioCVT{&cvt, nil}
	return rcvt, int(ret)
}

func ConvertAudio(cvt *AudioCVT) bool {
	//var buf2 [uint(cvt.cAudioCVT.len) * uint(cvt.cAudioCVT.len_mult)]byte
	//cvt.Buf = buf2
	//cvt.cAudioCVT.buf = (*C.Uint8)(unsafe.Pointer(&cvt.Buf[0]))
	//	cvt.cAudioCVT.len = C.int(len(buf))
	ret := C.SDL_ConvertAudio(cvt.cAudioCVT)
	return int(ret) == 0
}

func MixAudio(dst, src []byte, volume int) {
	C.SDL_MixAudio((*C.Uint8)(unsafe.Pointer(&dst[0])), (*C.Uint8)(unsafe.Pointer(&src[0])), C.Uint32(len(dst)), C.int(volume))
}

func MixAudioFormat(dst, src []byte, format AudioFormat, volume int) {
	C.SDL_MixAudioFormat((*C.Uint8)(unsafe.Pointer(&dst[0])), (*C.Uint8)(unsafe.Pointer(&src[0])), C.SDL_AudioFormat(int(format)), C.Uint32(len(dst)), C.int(volume))
}

func LockAudio() {
	C.SDL_LockAudio()
}

func UnlockAudio() {
	C.SDL_UnlockAudio()
}

func LockAudioDevice(dev AudioDeviceID) {
	C.SDL_LockAudioDevice(C.SDL_AudioDeviceID(int(dev)))
}

func UnlockAudioDevice(dev AudioDeviceID) {
	C.SDL_UnlockAudioDevice(C.SDL_AudioDeviceID(int(dev)))
}
