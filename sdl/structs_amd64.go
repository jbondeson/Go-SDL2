package sdl

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
	Alpha  uint8
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

type JoystickID int32

type JoyAxisEvent struct {
	Type       uint32
	Timestamp  uint32
	Which      JoystickID
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
	Which      JoystickID
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
	Which      JoystickID
	Hat        uint8
	Value      uint8
	Padding1   uint8
	Padding2   uint8
}

type JoyButtonEvent struct {
	Type       uint32
	Timestamp  uint32
	Which      JoystickID
	Button     uint8
	State      uint8
	Padding1   uint8
	Padding2   uint8
}

type JoyDeviceEvent struct {
	Type       uint32
	Timestamp  uint32
	Which      JoystickID
}

/**
 *  \brief Game controller axis motion event structure (event.caxis.*)
 */
type ControllerAxisEvent struct
{
	Type uint32        /**< ::CONTROLLERAXISMOTION */
	Timestamp uint32
	Which JoystickID /**< The joystick instance id */
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
	Which JoystickID /**< The joystick instance id */
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

type TouchID int32
type FingerID int32
type GestureID int32

/**
 *  \brief Touch finger event structure (event.tfinger.*)
 */
type TouchFingerEvent struct
{
	Type uint32        /**< ::FINGERMOTION or ::FINGERDOWN or ::FINGERUP */
	Timestamp uint32
	TouchId TouchID /**< The touch device id */
	FingerId FingerID
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
	TouchId TouchID /**< The touch device index */
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
	TouchId TouchID /**< The touch device id */
	GestureId GestureID
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

