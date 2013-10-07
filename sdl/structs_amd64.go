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

type PixelFormat struct {
	Format        uint32
	Palette       *Palette
	BitsPerPixel  uint8
	BytesPerPixel uint8
	Pad0          [2]byte
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	Refcount      int32
	Next          *PixelFormat
}

type Point struct {
	X int32
	Y int32
}

type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

type Color struct {
	R      uint8
	G      uint8
	B      uint8
	A      uint8
}

type Palette struct {
	Ncolors  int32
	Colors   *Color
	Version  uint32
	Refcount int32
}

type Scancode int32
type Keycode int32

type Keysym struct {
	Scancode int32
	Keycode  int32
	Mod      uint16
	Unused   uint32
}

// Fields shared by every event
type CommonEvent struct {
	Type uint32
	Timestamp uint32
}

type WindowEvent struct {
	Type uint32
	Timestamp uint32
	WindowID uint32
	Event uint8
	Padding1 uint8
	Padding2 uint8
	Padding3 uint8
	Data1 int32
	Data2 int32
}

type KeyboardEvent struct {
	Type uint32
	Timestamp uint32
	WindowID uint32
	State uint8
	Repeat uint8
	Padding2 uint8
	Padding3 uint8
	Keysym Keysym
}

type TextEditingEvent struct {
	Type uint32
	Timestamp uint32
	WindowID uint32
	Text [TEXTEDITINGEVENT_TEXT_SIZE]byte
	Start int32
	Length int32
}

type TextInputEvent struct {
	Type uint32
	Timestamp uint32
	WindowID uint32
	Text [TEXTEDITINGEVENT_TEXT_SIZE]byte
}

type MouseMotionEvent struct {
	Type        uint32
	Timestamp   uint32
	WindowID    uint32
	Which       uint32
	State       uint32
	X           int32
	Y           int32
	XRel        int32
	YRel        int32
}

type MouseButtonEvent struct {
	Type       uint32
	Timestamp  uint32
	WindowId   uint32
	Which      uint32
	Button     uint8
	State      uint8
	Padding1   uint8
	Padding2   uint8
	X          int32
	Y          int32
}

type MouseWheelEvent struct {
	Type       uint32
	Timestamp  uint32
	WindowID   uint32
	Which      uint32
	X          int32
	Y          int32
}

type JoyAxisEvent struct {
	Type       uint32
	Timestamp  uint32
	Which      int32
	Axis       uint8
	Padding1   uint8
	Padding2   uint8
	Padding3   uint8
	Value      int16
	Padding4   uint16
}

type JoyBallEvent struct {
	Type       uint32
	Timestamp  uint32
	Which      int32
	Ball       uint8
	Padding1   uint8
	Padding2   uint8
	Padding3   uint8
	Xrel       int16
	Yrel       int16
}

type JoyHatEvent struct {
	Type       uint32
	Timestamp  uint32
	Which      int32
	Hat        uint8
	Value      uint8
	Padding1   uint8
	Padding2   uint8
}

type JoyButtonEvent struct {
	Type       uint32
	Timestamp  uint32
	Which      int32
	Button     uint8
	State      uint8
	Padding1   uint8
	Padding2   uint8
}

type JoyDeviceEvent struct {
	Type       uint32
	Timestamp  uint32
	Which      int32
}

/**
 *  \brief Game controller axis motion event structure (event.caxis.*)
 */
type ControllerAxisEvent struct
{
	Type uint32        /**< ::CONTROLLERAXISMOTION */
	Timestamp uint32
	Which int32 /**< The joystick instance id */
	Axis uint8         /**< The controller axis (GameControllerAxis) */
	Padding1 uint8
	Padding2 uint8
	Padding3 uint8
	Value int16       /**< The axis value (range: -32768 to 32767) */
	Padding4 uint16
}


/**
 *  \brief Game controller button event structure (event.cbutton.*)
 */
type ControllerButtonEvent struct
{
	Type uint32        /**< ::CONTROLLERBUTTONDOWN or ::CONTROLLERBUTTONUP */
	Timestamp uint32
	Which int32 /**< The joystick instance id */
	Button uint8       /**< The controller button (GameControllerButton) */
	State uint8        /**< ::PRESSED or ::RELEASED */
	Padding1 uint8
	Padding2 uint8
}


/**
 *  \brief Controller device event structure (event.cdevice.*)
 */
type ControllerDeviceEvent struct
{
	Type uint32        /**< ::CONTROLLERDEVICEADDED, ::CONTROLLERDEVICEREMOVED, or ::CONTROLLERDEVICEREMAPPED */
	Timestamp uint32
	Which int32       /**< The joystick device index for the ADDED event, instance id for the REMOVED or REMAPPED event */
}

/**
 *  \brief Touch finger event structure (event.tfinger.*)
 */
type TouchFingerEvent struct
{
	Type uint32        /**< ::FINGERMOTION or ::FINGERDOWN or ::FINGERUP */
	Timestamp uint32
	TouchId int32 /**< The touch device id */
	FingerId int32
	X float32            /**< Normalized in the range 0...1 */
	Y float32            /**< Normalized in the range 0...1 */
	Dx float32           /**< Normalized in the range 0...1 */
	Dy float32           /**< Normalized in the range 0...1 */
	Pressure float32     /**< Normalized in the range 0...1 */
}


/**
 *  \brief Multiple Finger Gesture Event (event.mgesture.*)
 */
type MultiGestureEvent struct
{
	Type uint32        /**< ::MULTIGESTURE */
	Timestamp uint32
	TouchId int32 /**< The touch device index */
	Dtheta float32
	Ddist float32
	X float32
	Y float32
	NumFingers uint16
	Padding uint16
}

/**
 * \brief Dollar Gesture Event (event.dgesture.*)
 */
type DollarGestureEvent struct
{
	Type uint32        /**< ::DOLLARGESTURE */
	Timestamp uint32
	TouchId int32 /**< The touch device id */
	GestureId int32
	NumFingers uint32
	Error float32
	X float32            /**< Normalized center of gesture */
	Y float32            /**< Normalized center of gesture */
}

/**
 *  \brief An event used to request a file open by the system (event.drop.*)
 *         This event is disabled by default, you can enable it with SDL_EventState()
 *  \note If you enable this event, you must free the filename in the event.
 */
type DropEvent struct {
	Type       uint32 /**< ::DROPFILE */
	Timestamp  uint32
	File       *byte /**< The file name, which should be freed with SDL_free() */
}

type QuitEvent struct {
	Type uint32
	Timestamp uint32
}

type OSEvent struct {
	Type uint32
	Timestamp uint32
}

type UserEvent struct {
	Type  uint32
	Timestamp uint32
	WindowID uint32
	Code int32
	Data1 *byte
	Data2 *byte
}

type SysWMmsg struct{}

type SysWMEvent struct {
	Type uint32
	Timestamp uint32
	Msg  *SysWMmsg
}

type Event struct {
	Type uint32
	Timestamp uint32
	Pad0 [48]byte
}

