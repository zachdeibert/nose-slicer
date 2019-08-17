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
	case Microns, Millimeters, Centimeters, Meters:
		return gcode.UseMillimeters()
	case Thou, Inches:
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

// ApplyTo applies a unit to the options and gcode so all further commands can
// directly use what is in the options structure
func (unit Unit) ApplyTo(opts *Options, gcode gcode.Writer) error {
	multiplier, err := unit.GetMultiplier()
	if err != nil {
		return err
	}
	if err = unit.Use(gcode); err != nil {
		return err
	}
	opts.Diameter *= multiplier
	opts.Height *= multiplier
	opts.StartZ *= multiplier
	opts.EndZ *= multiplier
	opts.LayerHeight *= multiplier
	opts.FeedRate *= multiplier
	opts.RapidFeedRate *= multiplier
	opts.MaterialDiameter *= multiplier
	opts.EndMillDiameter *= multiplier
	return nil
}
