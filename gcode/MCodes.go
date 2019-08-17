package gcode

import "io"

// SpindleOn turns on the spindle
func (writer *Writer) SpindleOn() error {
	_, err := io.WriteString(writer.Stream, "M3\n")
	return err
}

// SpindleOff turns off the spindle
func (writer *Writer) SpindleOff() error {
	_, err := io.WriteString(writer.Stream, "M5\n")
	return err
}
