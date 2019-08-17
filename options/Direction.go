package options

import (
	"errors"

	"../gcode"
)

// Direction represents the direction the cut should be made
type Direction int

const (
	// Clockwise is the way the clock goes
	Clockwise Direction = 1
	// Counterclockwise isn't the way the clock goes
	Counterclockwise Direction = 2
)

// MakeCircle generates the gcode to make a complete circle in the given direction
func (dir Direction) MakeCircle(radius float64, feedRate float64, currentX float64, currentY float64, gcode gcode.Writer) error {
	if currentX != 0 && currentY != 0 {
		return errors.New("Can only start on one of the axes")
	} else if currentX == 0 && currentY == 0 {
		return errors.New("Cannot start at the origin")
	} else if currentX != radius && currentY != radius {
		return errors.New("Has to start at a point on the circle")
	} else if radius <= 0 {
		return errors.New("Radius must be positive")
	}
	switch dir {
	case Clockwise:
		for i := 0; i < 4; i++ {
			if currentX == 0 {
				if currentY < 0 {
					currentX = -radius
					currentY = 0
				} else {
					currentX = radius
					currentY = 0
				}
			} else if currentX < 0 {
				currentX = 0
				currentY = radius
			} else {
				currentX = 0
				currentY = -radius
			}
			if err := gcode.ClockwiseInterpolation(currentX, currentY, radius, feedRate); err != nil {
				return err
			}
		}
		return nil
	case Counterclockwise:
		for i := 0; i < 4; i++ {
			if currentX == 0 {
				if currentY < 0 {
					currentX = radius
					currentY = 0
				} else {
					currentX = -radius
					currentY = 0
				}
			} else if currentX < 0 {
				currentX = 0
				currentY = -radius
			} else {
				currentX = 0
				currentY = radius
			}
			if err := gcode.CounterclockwiseInterpolation(currentX, currentY, radius, feedRate); err != nil {
				return err
			}
		}
		return nil
	default:
		return errors.New("Invalid constant")
	}
}
