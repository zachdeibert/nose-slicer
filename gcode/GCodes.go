package gcode

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strings"
)

// RapidPositioning moves the toolhead to a new position as fast as possible
func (writer Writer) RapidPositioning(x, y, z float64) error {
	var cmd strings.Builder
	cmd.WriteString("G0")
	if !math.IsNaN(x) && !math.IsInf(x, 0) {
		fmt.Fprintf(&cmd, " X%f", x)
	}
	if !math.IsNaN(y) && !math.IsInf(y, 0) {
		fmt.Fprintf(&cmd, " Y%f", y)
	}
	if !math.IsNaN(z) && !math.IsInf(z, 0) {
		fmt.Fprintf(&cmd, " Z%f", z)
	}
	cmd.WriteString("\n")
	_, err := io.WriteString(writer.Stream, cmd.String())
	return err
}

// LinearInterpolation moves the toolhead at a specific speeed to a new position
func (writer Writer) LinearInterpolation(x, y, z, feed float64) error {
	if math.IsNaN(feed) || math.IsInf(feed, 0) || feed <= 0 {
		return errors.New("Invalid feedrate")
	}
	var cmd strings.Builder
	cmd.WriteString("G1")
	if !math.IsNaN(x) && !math.IsInf(x, 0) {
		fmt.Fprintf(&cmd, " X%f", x)
	}
	if !math.IsNaN(y) && !math.IsInf(y, 0) {
		fmt.Fprintf(&cmd, " Y%f", y)
	}
	if !math.IsNaN(z) && !math.IsInf(z, 0) {
		fmt.Fprintf(&cmd, " Z%f", z)
	}
	fmt.Fprintf(&cmd, " F%f\n", feed)
	_, err := io.WriteString(writer.Stream, cmd.String())
	return err
}

// ClockwiseInterpolation moves the toolhead at a specific speeed to a new
// position around a clockwise arc
func (writer Writer) ClockwiseInterpolation(x, y, radius, feed float64) error {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return errors.New("Invalid x")
	}
	if math.IsNaN(y) || math.IsInf(y, 0) {
		return errors.New("Invalid y")
	}
	if math.IsNaN(radius) || math.IsInf(radius, 0) || radius <= 0 {
		return errors.New("Invalid radius")
	}
	if math.IsNaN(feed) || math.IsInf(feed, 0) || feed <= 0 {
		return errors.New("Invalid feedrate")
	}
	_, err := fmt.Fprintf(writer.Stream, "G2 X%f Y%f R%f F%f\n", x, y, radius, feed)
	return err
}

// CounterclockwiseInterpolation moves the toolhead at a specific speeed to a
// new position around a clockwise arc
func (writer Writer) CounterclockwiseInterpolation(x, y, radius, feed float64) error {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return errors.New("Invalid x")
	}
	if math.IsNaN(y) || math.IsInf(y, 0) {
		return errors.New("Invalid y")
	}
	if math.IsNaN(radius) || math.IsInf(radius, 0) || radius <= 0 {
		return errors.New("Invalid radius")
	}
	if math.IsNaN(feed) || math.IsInf(feed, 0) || feed <= 0 {
		return errors.New("Invalid feedrate")
	}
	_, err := fmt.Fprintf(writer.Stream, "G3 X%f Y%f R%f F%f\n", x, y, radius, feed)
	return err
}

// Dwell waits for a specified amount of time
func (writer Writer) Dwell(time float64) error {
	if math.IsNaN(time) || math.IsInf(time, 0) || time <= 0 {
		return errors.New("Invalid time")
	}
	_, err := fmt.Fprintf(writer.Stream, "G4 P%f\n", time)
	return err
}
