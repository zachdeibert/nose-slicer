package gcode

import "io"

// UseInches converts the machine into inch mode (why would anyone ever want
// that?)
func (writer Writer) UseInches() error {
	_, err := io.WriteString(writer.Stream, "G20\n")
	return err
}

// UseMillimeters converts the machine into millimeter mode
func (writer Writer) UseMillimeters() error {
	_, err := io.WriteString(writer.Stream, "G21\n")
	return err
}

// UseAbsoluteProgramming makes all coordinates absolute
func (writer Writer) UseAbsoluteProgramming() error {
	_, err := io.WriteString(writer.Stream, "G90\n")
	return err
}

// UseRelativeProgramming makes all coordinates relative
func (writer Writer) UseRelativeProgramming() error {
	_, err := io.WriteString(writer.Stream, "G91\n")
	return err
}
