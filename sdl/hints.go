package sdl

/*
  SDL Go Wrapper

  Simple DirectMedia Layer
  Copyright (C) 1997-2014 Jeremy Bondeson <jbondeson@gmail.com>
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
	/**
	 *  \brief  A variable controlling how 3D acceleration is used to accelerate the SDL screen surface.
	 *
	 *  SDL can try to accelerate the SDL screen surface by using streaming
	 *  textures with a 3D rendering engine.  This variable controls whether and
	 *  how this is done.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Disable 3D acceleration
	 *    "1"       - Enable 3D acceleration, using the default renderer.
	 *    "X"       - Enable 3D acceleration, using X where X is one of the valid rendering drivers.  (e.g. "direct3d", "opengl", etc.)
	 *
	 *  By default SDL tries to make a best guess for each platform whether
	 *  to use acceleration or not.
	 */
	HINT_FRAMEBUFFER_ACCELERATION = "SDL_FRAMEBUFFER_ACCELERATION"

	/**
	 *  \brief  A variable specifying which render driver to use.
	 *
	 *  If the application doesn't pick a specific renderer to use, this variable
	 *  specifies the name of the preferred renderer.  If the preferred renderer
	 *  can't be initialized, the normal default renderer is used.
	 *
	 *  This variable is case insensitive and can be set to the following values:
	 *    "direct3d"
	 *    "opengl"
	 *    "opengles2"
	 *    "opengles"
	 *    "software"
	 *
	 *  The default varies by platform, but it's the first one in the list that
	 *  is available on the current platform.
	 */
	HINT_RENDER_DRIVER = "SDL_RENDER_DRIVER"

	/**
	 *  \brief  A variable controlling whether the OpenGL render driver uses shaders if they are available.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Disable shaders
	 *    "1"       - Enable shaders
	 *
	 *  By default shaders are used if OpenGL supports them.
	 */
	HINT_RENDER_OPENGL_SHADERS = "SDL_RENDER_OPENGL_SHADERS"

	/**
	 *  \brief  A variable controlling whether the Direct3D device is initialized for thread-safe operations.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Thread-safety is not enabled (faster)
	 *    "1"       - Thread-safety is enabled
	 *
	 *  By default the Direct3D device is created with thread-safety disabled.
	 */
	HINT_RENDER_DIRECT3D_THREADSAFE = "SDL_RENDER_DIRECT3D_THREADSAFE"

	/**
	 *  \brief  A variable controlling whether to enable Direct3D 11+'s Debug Layer.
	 *
	 *  This variable does not have any effect on the Direct3D 9 based renderer.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Disable Debug Layer use
	 *    "1"       - Enable Debug Layer use
	 *
	 *  By default, SDL does not use Direct3D Debug Layer.
	 */
	HINT_RENDER_DIRECT3D11_DEBUG = "SDL_HINT_RENDER_DIRECT3D11_DEBUG"

	/**
	 *  \brief  A variable controlling the scaling quality
	 *
	 *  This variable can be set to the following values:
	 *    "0" or "nearest" - Nearest pixel sampling
	 *    "1" or "linear"  - Linear filtering (supported by OpenGL and Direct3D)
	 *    "2" or "best"    - Currently this is the same as "linear"
	 *
	 *  By default nearest pixel sampling is used
	 */
	HINT_RENDER_SCALE_QUALITY = "SDL_RENDER_SCALE_QUALITY"
	/**
	 *  \brief  A variable controlling whether updates to the SDL screen surface should be synchronized with the vertical refresh, to avoid tearing.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Disable vsync
	 *    "1"       - Enable vsync
	 *
	 *  By default SDL does not sync screen surface updates with vertical refresh.
	 */
	HINT_RENDER_VSYNC = "SDL_RENDER_VSYNC"

	/**
	 *  \brief  A variable controlling whether the screensaver is enabled.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Disable screensaver
	 *    "1"       - Enable screensaver
	 *
	 *  By default SDL will disable the screensaver.
	 */
	HINT_VIDEO_ALLOW_SCREENSAVER = "SDL_VIDEO_ALLOW_SCREENSAVER"

	/**
	 *  \brief  A variable controlling whether the X11 VidMode extension should be used.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Disable XVidMode
	 *    "1"       - Enable XVidMode
	 *
	 *  By default SDL will use XVidMode if it is available.
	 */
	HINT_VIDEO_X11_XVIDMODE = "SDL_VIDEO_X11_XVIDMODE"

	/**
	 *  \brief  A variable controlling whether the X11 Xinerama extension should be used.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Disable Xinerama
	 *    "1"       - Enable Xinerama
	 *
	 *  By default SDL will use Xinerama if it is available.
	 */
	HINT_VIDEO_X11_XINERAMA = "SDL_VIDEO_X11_XINERAMA"

	/**
	 *  \brief  A variable controlling whether the X11 XRandR extension should be used.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Disable XRandR
	 *    "1"       - Enable XRandR
	 *
	 *  By default SDL will not use XRandR because of window manager issues.
	 */
	HINT_VIDEO_X11_XRANDR = "SDL_VIDEO_X11_XRANDR"

	/**
	 *  \brief  A variable controlling whether grabbing input grabs the keyboard
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Grab will affect only the mouse
	 *    "1"       - Grab will affect mouse and keyboard
	 *
	 *  By default SDL will not grab the keyboard so system shortcuts still work.
	 */
	HINT_GRAB_KEYBOARD = "SDL_GRAB_KEYBOARD"

	/**
	 *  \brief  A variable controlling whether relative mouse mode is implemented using mouse warping
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Relative mouse mode uses raw input
	 *    "1"       - Relative mouse mode uses mouse warping
	 *
	 *  By default SDL will use raw input for relative mouse mode
	 */
	HINT_MOUSE_RELATIVE_MODE_WARP = "SDL_MOUSE_RELATIVE_MODE_WARP"

	/**
	 *  \brief Minimize your SDL_Window if it loses key focus when in fullscreen mode. Defaults to true.
	 *
	 */
	HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS = "SDL_VIDEO_MINIMIZE_ON_FOCUS_LOSS"

	/**
	 *  \brief  A variable controlling whether the idle timer is disabled on iOS.
	 *
	 *  When an iOS app does not receive touches for some time, the screen is
	 *  dimmed automatically. For games where the accelerometer is the only input
	 *  this is problematic. This functionality can be disabled by setting this
	 *  hint.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - Enable idle timer
	 *    "1"       - Disable idle timer
	 */
	HINT_IDLE_TIMER_DISABLED = "SDL_IOS_IDLE_TIMER_DISABLED"

	/**
	 *  \brief  A variable controlling which orientations are allowed on iOS.
	 *
	 *  In some circumstances it is necessary to be able to explicitly control
	 *  which UI orientations are allowed.
	 *
	 *  This variable is a space delimited list of the following values:
	 *    "LandscapeLeft", "LandscapeRight", "Portrait" "PortraitUpsideDown"
	 */
	HINT_ORIENTATIONS = "SDL_IOS_ORIENTATIONS"

	/**
	 *  \brief  A variable controlling whether an Android built-in accelerometer should be
	 *  listed as a joystick device, rather than listing actual joysticks only.
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - List only real joysticks and accept input from them
	 *    "1"       - List real joysticks along with the accelerometer as if it were a 3 axis joystick (the default).
	 */
	HINT_ACCELEROMETER_AS_JOYSTICK = "SDL_ACCELEROMETER_AS_JOYSTICK"

	/**
	 *  \brief  A variable that lets you disable the detection and use of Xinput gamepad devices
	 *
	 *  The variable can be set to the following values:
	 *    "0"       - Disable XInput detection (only uses direct input)
	 *    "1"       - Enable XInput detection (the default)
	 */
	HINT_XINPUT_ENABLED = "SDL_XINPUT_ENABLED"

	/**
	 *  \brief  A variable that lets you manually hint extra gamecontroller db entries
	 *
	 *  The variable should be newline delimited rows of gamecontroller config data, see SDL_gamecontroller.h
	 *
	 *  This hint must be set before calling SDL_Init(SDL_INIT_GAMECONTROLLER)
	 *  You can update mappings after the system is initialized with SDL_GameControllerMappingForGUID() and SDL_GameControllerAddMapping()
	 */
	HINT_GAMECONTROLLERCONFIG = "SDL_GAMECONTROLLERCONFIG"

	/**
	 *  \brief  A variable that lets you enable joystick (and gamecontroller) events even when your app is in the background.
	 *
	 *  The variable can be set to the following values:
	 *    "0"       - Disable joystick & gamecontroller input events when the
	 *                application is in the background.
	 *    "1"       - Enable joystick & gamecontroller input events when the
	 *                application is in the background.
	 *
	 *  The default value is "0".  This hint may be set at any time.
	 */
	HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS = "SDL_JOYSTICK_ALLOW_BACKGROUND_EVENTS"

	/**
	 *  \brief If set to 0 then never set the top most bit on a SDL Window, even if the video mode expects it.
	 *      This is a debugging aid for developers and not expected to be used by end users. The default is "1"
	 *
	 *  This variable can be set to the following values:
	 *    "0"       - don't allow topmost
	 *    "1"       - allow topmost
	 */
	HINT_ALLOW_TOPMOST = "SDL_ALLOW_TOPMOST"

	/**
	 *  \brief A variable that controls the timer resolution, in milliseconds.
	 *
	 *  The higher resolution the timer, the more frequently the CPU services
	 *  timer interrupts, and the more precise delays are, but this takes up
	 *  power and CPU time.  This hint is only used on Windows 7 and earlier.
	 *
	 *  See this blog post for more information:
	 *  http://randomascii.wordpress.com/2013/07/08/windows-timer-resolution-megawatts-wasted/
	 *
	 *  If this variable is set to "0", the system timer resolution is not set.
	 *
	 *  The default value is "1". This hint may be set at any time.
	 */
	HINT_TIMER_RESOLUTION = "SDL_TIMER_RESOLUTION"

	/**
	 *  \brief If set to 1, then do not allow high-DPI windows. ("Retina" on Mac)
	 */
	HINT_VIDEO_HIGHDPI_DISABLED = "SDL_VIDEO_HIGHDPI_DISABLED"

	/**
	 *  \brief A variable that determines whether ctrl+click should generate a right-click event on Mac
	 *
	 *  If present, holding ctrl while left clicking will generate a right click
	 *  event when on Mac.
	 */
	HINT_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK = "SDL_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK"

	/**
	 *  \brief  A variable specifying which shader compiler to preload when using the Chrome ANGLE binaries
	 *
	 *  SDL has EGL and OpenGL ES2 support on Windows via the ANGLE project. It
	 *  can use two different sets of binaries, those compiled by the user from source
	 *  or those provided by the Chrome browser. In the later case, these binaries require
	 *  that SDL loads a DLL providing the shader compiler.
	 *
	 *  This variable can be set to the following values:
	 *    "d3dcompiler_46.dll" - default, best for Vista or later.
	 *    "d3dcompiler_43.dll" - for XP support.
	 *    "none" - do not load any library, useful if you compiled ANGLE from source and included the compiler in your binaries.
	 *
	 */
	HINT_VIDEO_WIN_D3DCOMPILER = "SDL_VIDEO_WIN_D3DCOMPILER"

	/**
	 *  \brief  A variable that is the address of another SDL_Window* (as a hex string formatted with "%p").
	 *
	 *  If this hint is set before SDL_CreateWindowFrom() and the SDL_Window* it is set to has
	 *  SDL_WINDOW_OPENGL set (and running on WGL only, currently), then two things will occur on the newly
	 *  created SDL_Window:
	 *
	 *  1. Its pixel format will be set to the same pixel format as this SDL_Window.  This is
	 *  needed for example when sharing an OpenGL context across multiple windows.
	 *
	 *  2. The flag SDL_WINDOW_OPENGL will be set on the new window so it can be used for
	 *  OpenGL rendering.
	 *
	 *  This variable can be set to the following values:
	 *    The address (as a string "%p") of the SDL_Window* that new windows created with SDL_CreateWindowFrom() should
	 *    share a pixel format with.
	 */
	HINT_VIDEO_WINDOW_SHARE_PIXEL_FORMAT = "SDL_VIDEO_WINDOW_SHARE_PIXEL_FORMAT"

	/*
	 *  \brief A URL to a WinRT app's privacy policy
	 *
	 *  All network-enabled WinRT apps must make a privacy policy available to its
	 *  users.  On Windows 8, 8.1, and RT, Microsoft mandates that this policy be
	 *  be available in the Windows Settings charm, as accessed from within the app.
	 *  SDL provides code to add a URL-based link there, which can point to the app's
	 *  privacy policy.
	 *
	 *  To setup a URL to an app's privacy policy, set SDL_HINT_WINRT_PRIVACY_POLICY_URL
	 *  before calling any SDL_Init functions.  The contents of the hint should
	 *  be a valid URL.  For example, "http://www.example.com".
	 *
	 *  The default value is "", which will prevent SDL from adding a privacy policy
	 *  link to the Settings charm.  This hint should only be set during app init.
	 *
	 *  The label text of an app's "Privacy Policy" link may be customized via another
	 *  hint, SDL_HINT_WINRT_PRIVACY_POLICY_LABEL.
	 *
	 *  Please note that on Windows Phone, Microsoft does not provide standard UI
	 *  for displaying a privacy policy link, and as such, SDL_HINT_WINRT_PRIVACY_POLICY_URL
	 *  will not get used on that platform.  Network-enabled phone apps should display
	 *  their privacy policy through some other, in-app means.
	 */
	HINT_WINRT_PRIVACY_POLICY_URL = "SDL_HINT_WINRT_PRIVACY_POLICY_URL"

	/** \brief Label text for a WinRT app's privacy policy link
	 *
	 *  Network-enabled WinRT apps must include a privacy policy.  On Windows 8, 8.1, and RT,
	 *  Microsoft mandates that this policy be available via the Windows Settings charm.
	 *  SDL provides code to add a link there, with it's label text being set via the
	 *  optional hint, SDL_HINT_WINRT_PRIVACY_POLICY_LABEL.
	 *
	 *  Please note that a privacy policy's contents are not set via this hint.  A separate
	 *  hint, SDL_HINT_WINRT_PRIVACY_POLICY_URL, is used to link to the actual text of the
	 *  policy.
	 *
	 *  The contents of this hint should be encoded as a UTF8 string.
	 *
	 *  The default value is "Privacy Policy".  This hint should only be set during app
	 *  initialization, preferably before any calls to SDL_Init.
	 *
	 *  For additional information on linking to a privacy policy, see the documentation for
	 *  SDL_HINT_WINRT_PRIVACY_POLICY_URL.
	 */
	HINT_WINRT_PRIVACY_POLICY_LABEL = "SDL_HINT_WINRT_PRIVACY_POLICY_LABEL"

	/** \brief If set to 1, back button press events on Windows Phone 8+ will be marked as handled.
	 *
	 *  TODO, WinRT: document SDL_HINT_WINRT_HANDLE_BACK_BUTTON need and use
	 *  For now, more details on why this is needed can be found at the
	 *  beginning of the following web page:
	 *  http://msdn.microsoft.com/en-us/library/windowsphone/develop/jj247550(v=vs.105).aspx
	 */
	HINT_WINRT_HANDLE_BACK_BUTTON = "SDL_HINT_WINRT_HANDLE_BACK_BUTTON"

	/**
	 *  \brief  A variable that dictates policy for fullscreen Spaces on Mac OS X.
	 *
	 *  This hint only applies to Mac OS X.
	 *
	 *  The variable can be set to the following values:
	 *    "0"       - Disable Spaces support (FULLSCREEN_DESKTOP won't use them and
	 *                SDL_WINDOW_RESIZABLE windows won't offer the "fullscreen"
	 *                button on their titlebars).
	 *    "1"       - Enable Spaces support (FULLSCREEN_DESKTOP will use them and
	 *                SDL_WINDOW_RESIZABLE windows will offer the "fullscreen"
	 *                button on their titlebars.
	 *
	 *  The default value is "1". Spaces are disabled regardless of this hint if
	 *   the OS isn't at least Mac OS X Lion (10.7). This hint must be set before
	 *   any windows are created.
	 */
	HINT_VIDEO_MAC_FULLSCREEN_SPACES = "SDL_VIDEO_MAC_FULLSCREEN_SPACES"
)

const (
	HINT_DEFAULT  = C.SDL_HINT_DEFAULT
	HINT_NORMAL   = C.SDL_HINT_NORMAL
	HIST_OVERRIDE = C.SDL_HINT_OVERRIDE
)

func SetHint(name string, value string) bool {
	cname := C.CString(name)
	cvalue := C.CString(value)

	ret := C.SDL_SetHint(cname, cvalue)

	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(cvalue))

	return ret == C.SDL_TRUE
}

func SetHintWithPriority(name string, value string, priority int) bool {
	cname := C.CString(name)
	cvalue := C.CString(value)

	ret := C.SDL_SetHintWithPriority(cname, cvalue, C.SDL_HintPriority(priority))

	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(cvalue))

	return ret == C.SDL_TRUE
}

func GetHint(name string) string {
	cname := C.CString(name)

	cvalue := C.SDL_GetHint(cname)

	value := C.GoString(cvalue)

	C.free(unsafe.Pointer(cname))

	return value
}

//TODO: SDL_AddHintCallback
//TODO: SDL_DelHintCallback

func ClearHints() {
	C.SDL_ClearHints()
}