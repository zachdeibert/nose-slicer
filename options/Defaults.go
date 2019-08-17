package options

import "math"

// DefaultOptions returns all of the defaults
func DefaultOptions() Options {
	return Options{
		Mode:               Normal,
		Profile:            0,
		Diameter:           math.NaN(),
		Height:             math.NaN(),
		StartZ:             0,
		EndZ:               math.NaN(),
		LayerHeight:        0.5,
		FeedRate:           1000,
		RapidFeedRate:      0,
		LayerDwell:         0,
		MaterialDiameter:   math.NaN(),
		EndMillDiameter:    math.NaN(),
		Unit:               Millimeters,
		Direction:          Clockwise,
		PassOverlapPercent: 0.3,
		FlatLayers:         2,
		OutputFile:         nil,
	}
}
