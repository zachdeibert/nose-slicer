package options

import (
	"errors"
	"math"

	"../gcode"
)

// Unit describes what unit the measurements are in
type Unit int

const (
	// Microns is 1/1000 of a millimeter
	Microns Unit = 1
	// Millimeters are millimeters
	Millimeters Unit = 2
	// Centimeters are 10 millimeters
	Centimeters Unit = 3
	// Meters are 1000 millimeters
	Meters Unit = 4
	// Thou are 1/1000 of an inch
	Thou Unit = 5
	// Inches are gross
	Inches Unit = 6
)

// Use writes the gcodes to set the unit type
func (unit Unit) Use(gcode gcode.Writer) error {
	switch unit {
	case Microns:
		return gcode.UseMillimeters()
	case Millimeters:
		return gcode.UseMillimeters()
	case Centimeters:
		return gcode.UseMillimeters()
	case Meters:
		return gcode.UseMillimeters()
	case Thou:
		return gcode.UseInches()
	case Inches:
		return gcode.UseInches()
	default:
		return errors.New("Invalid constant")
	}
}

// GetMultiplier gets a constant that all options should be multiplied by to be
// compatible with the CNC
func (unit Unit) GetMultiplier() (float64, error) {
	switch unit {
	case Microns:
		return 0.001, nil
	case Millimeters:
		return 1, nil
	case Centimeters:
		return 10, nil
	case Meters:
		return 1000, nil
	case Thou:
		return 0.001, nil
	case Inches:
		return 1, nil
	default:
		return math.NaN(), errors.New("Invalid constant")
	}
}
