package input

import (
	"log"
	"mousephoneserver/internal/command"
	"syscall"
	"unsafe"
)

// Load user32 Windows dll
var (
	user32           = syscall.NewLazyDLL("user32.dll")
	procSetCursorPos = user32.NewProc("SetCursorPos")
	procMouseEvent   = user32.NewProc("mouse_event")
	procGetCursorPos = user32.NewProc("GetCursorPos")
)

// Mouse event constants
const (
	MOUSEEVENTF_LEFTDOWN  = 0x0002
	MOUSEEVENTF_LEFTUP    = 0x0004
	MOUSEEVENTF_RIGHTDOWN = 0x0008
	MOUSEEVENTF_RIGHTUP   = 0x0010
)

type Point struct {
	X, Y int32
}

// Controller Object that executes the commands
type Controller struct{}

// NewController creates and returns new instance of controller
func NewController() *Controller {
	return &Controller{}
}

// Execute receives a command and execute it in the OS. Controller's method
func (c *Controller) Execute(cmd *command.Command) {
	switch cmd.Type {
	case "move":
		var p Point

		// Current position of mouse
		procGetCursorPos.Call(uintptr(unsafe.Pointer(&p)))

		// Calculate new position
		newX := p.X + int32(cmd.Dx)
		newY := p.Y + int32(cmd.Dy)

		// Move mouse to new position
		procSetCursorPos.Call(uintptr(newX), uintptr(newY))

	case "click":
		switch cmd.Button {
		case "left":
			// Simula o botão esquerdo sendo pressionado e depois solto
			procMouseEvent.Call(MOUSEEVENTF_LEFTDOWN, 0, 0, 0, 0)
			procMouseEvent.Call(MOUSEEVENTF_LEFTUP, 0, 0, 0, 0)
		case "right":
			// Simula o botão direito sendo pressionado e depois solto
			procMouseEvent.Call(MOUSEEVENTF_RIGHTDOWN, 0, 0, 0, 0)
			procMouseEvent.Call(MOUSEEVENTF_RIGHTUP, 0, 0, 0, 0)
		}

	default:
		log.Printf("Unknown command received: %s", cmd.Type)
	}
}
