package gcode

import (
	"io"
	"math"
)

// Writer is an object that can write GCode to a stream
type Writer struct {
	Stream io.Writer
	X      float64
	Y      float64
	Z      float64
	Time   float64
}

// CreateWriter makes a new GCode writer
func CreateWriter(stream io.Writer) Writer {
	return Writer{
		Stream: stream,
		X:      math.NaN(),
		Y:      math.NaN(),
		Z:      math.NaN(),
		Time:   0,
	}
}
