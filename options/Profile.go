package options

import (
	"errors"
	"math"
)

// Profile represents a cross-section type for the nose cone
type Profile int

const (
	// Conical is a nose cone that looks like a cone
	Conical Profile = 1
	// Ogive is a nose cone constructed from a circle
	Ogive Profile = 2
)

// GetRadius calculates the radius of a layer of the nose cone
func (profile Profile) GetRadius(pos, diameter, height float64) (float64, error) {
	if math.IsNaN(pos) || math.IsInf(pos, 0) || pos < 0 || pos > height {
		return math.NaN(), errors.New("Invalid position")
	} else if math.IsNaN(diameter) || math.IsInf(diameter, 0) || diameter <= 0 {
		return math.NaN(), errors.New("Invalid diameter")
	} else if math.IsNaN(height) || math.IsInf(height, 0) || height <= 0 {
		return math.NaN(), errors.New("Invalid height")
	}
	switch profile {
	case Conical:
		return diameter / 2 * pos / height, nil
	case Ogive:
		fullAngle := 2 * math.Asin(math.Sqrt(height*height+diameter*diameter)/2)
		radius := diameter / math.Sin(fullAngle)
		angle := math.Asin((height - pos) / radius)
		return diameter/2 - radius*(1-math.Cos(angle)), nil
	default:
		return math.NaN(), errors.New("Invalid constant")
	}
}
