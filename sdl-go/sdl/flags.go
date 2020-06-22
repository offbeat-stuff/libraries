package sdl

/*
#include "main.h"

#cgo amd64 CFLAGS: -I include
#cgo LDFLAGS: -lSDL2
*/
import "C"

const (
	//InitTimer |-> timer subsystem
	InitTimer = C.SDL_INIT_TIMER
	//InitAudio |-> audio subsystem
	InitAudio = C.SDL_INIT_AUDIO
	//InitVideo |-> video subsystem; automatically initializes the events subsystem
	InitVideo = C.SDL_INIT_VIDEO
	//InitJoystick |-> joystick subsystem; automatically initializes the events subsystem
	InitJoystick = C.SDL_INIT_JOYSTICK
	//InitHaptic |-> haptic (force feedback) subsystem
	InitHaptic = C.SDL_INIT_HAPTIC
	//InitGameController |->  controller subsystem; automatically initializes the joystick subsystem
	InitGameController = C.SDL_INIT_GAMECONTROLLER
	//InitEvents |-> events subsystem
	InitEvents = C.SDL_INIT_EVENTS
	//InitNoParachute |-> compatibility; this flag is ignored
	InitNoParachute = C.SDL_INIT_NOPARACHUTE
	//InitSensor |-> sensor subsystem
	InitSensor = C.SDL_INIT_SENSOR
	//InitEverything |-> all of the above subsystems
	InitEverything = C.SDL_INIT_EVERYTHING
)

const (
	//WindowFullscreen |-> fullscreen window
	WindowFullscreen = C.SDL_WINDOW_FULLSCREEN
	//WindowOpenGL |-> window usable with OpenGL context
	WindowOpenGL = C.SDL_WINDOW_OPENGL
	//WindowShown |-> window is visible
	WindowShown = C.SDL_WINDOW_SHOWN
	//WindowHidden |-> window is not visible
	WindowHidden = C.SDL_WINDOW_HIDDEN
	//WindowBorderless |-> no window decoration
	WindowBorderless = C.SDL_WINDOW_BORDERLESS
	//WindowResizable |-> window can be resized
	WindowResizable = C.SDL_WINDOW_RESIZABLE
	//WindowMinimized |-> window is minimized
	WindowMinimized = C.SDL_WINDOW_MINIMIZED
	//WindowMaximized |-> window is maximized
	WindowMaximized = C.SDL_WINDOW_MAXIMIZED
	//WindowInputGrabbed |-> window has grabbed input focus
	WindowInputGrabbed = C.SDL_WINDOW_INPUT_GRABBED
	//WindowInputFocus |-> window has input focus
	WindowInputFocus = C.SDL_WINDOW_INPUT_FOCUS
	//WindowMouseFocus |-> window has mouse focus
	WindowMouseFocus = C.SDL_WINDOW_MOUSE_FOCUS
	//WindowFullscrennDesktop |-> fullscreen window at the current desktop resolution
	WindowFullscrennDesktop = C.SDL_WINDOW_FULLSCREEN_DESKTOP
	//WindowForeign |-> window not created by SDL
	WindowForeign = C.SDL_WINDOW_FOREIGN
	//WindowAllowHighDPI |-> window should be created in high-DPI mode if supported (>= SDL 2.0.1)
	WindowAllowHighDPI = C.SDL_WINDOW_ALLOW_HIGHDPI
	//WindowMouseCapture |-> window has mouse captured (unrelated to INPUT_GRABBED, >= SDL 2.0.4)
	WindowMouseCapture = C.SDL_WINDOW_MOUSE_CAPTURE
	//WindowAlwaysOnTop |-> window should always be above others (X11 only, >= SDL 2.0.5)
	WindowAlwaysOnTop = C.SDL_WINDOW_ALWAYS_ON_TOP
	//WindowSkipTaskbar |-> window should not be added to the taskbar (X11 only, >= SDL 2.0.5)
	WindowSkipTaskbar = C.SDL_WINDOW_SKIP_TASKBAR
	//WindowUtility |-> window should be treated as a utility window (X11 only, >= SDL 2.0.5)
	WindowUtility = C.SDL_WINDOW_UTILITY
	//WindowTooltip |-> window should be treated as a tooltip (X11 only, >= SDL 2.0.5)
	WindowTooltip = C.SDL_WINDOW_TOOLTIP
	//WindowPopupMenu |-> window should be treated as a popup menu (X11 only, >= SDL 2.0.5)
	WindowPopupMenu = C.SDL_WINDOW_POPUP_MENU
	//WindowVulkan |-> window usable for Vulkan surface (>= SDL 2.0.6)
	WindowVulkan = C.SDL_WINDOW_VULKAN
)
