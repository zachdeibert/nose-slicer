package gcode

import (
	"io"
)

// Writer is an object that can write GCode to a stream
type Writer struct {
	Stream io.Writer
}
