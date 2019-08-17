package logic

import (
	"math"

	"../gcode"
	"../options"
)

func move(opts options.Options, gcode *gcode.Writer, x, y, z float64) error {
	if opts.RapidFeedRate == 0 {
		return gcode.RapidPositioning(x, y, z)
	} else {
		return gcode.LinearInterpolation(x, y, z, opts.RapidFeedRate)
	}
}

func millLayer(opts options.Options, gcode *gcode.Writer, layerWidth, outerLayer, offModel, z, insideRadius float64) error {
	if err := move(opts, gcode, outerLayer, 0, z); err != nil {
		return err
	}
	stop := insideRadius + layerWidth*3/2
	for x := outerLayer; x > stop; x -= layerWidth {
		if err := gcode.LinearInterpolation(x, 0, math.NaN(), opts.FeedRate); err != nil {
			return err
		}
		if err := opts.Direction.MakeCircle(x, opts.FeedRate, x, 0, gcode); err != nil {
			return err
		}
	}
	stop -= layerWidth
	if err := gcode.LinearInterpolation(stop, 0, math.NaN(), opts.FeedRate); err != nil {
		return err
	}
	if err := opts.Direction.MakeCircle(stop, opts.FeedRate, stop, 0, gcode); err != nil {
		return err
	}
	if err := move(opts, gcode, outerLayer, 0, math.NaN()); err != nil {
		return err
	}
	if opts.LayerDwell > 0 {
		if err := move(opts, gcode, offModel, 0, math.NaN()); err != nil {
			return err
		}
		if err := gcode.SpindleOff(); err != nil {
			return err
		}
		if err := gcode.Dwell(opts.LayerDwell); err != nil {
			return err
		}
		if err := gcode.SpindleOn(); err != nil {
			return err
		}
	}
	return nil
}

// Slice generates the GCode to make the nose cone
func Slice(opts options.Options, gcode *gcode.Writer) error {
	if err := gcode.UseAbsoluteProgramming(); err != nil {
		return err
	}
	if err := opts.Unit.ApplyTo(&opts, gcode); err != nil {
		return err
	}
	layerWidth := (1 - opts.PassOverlapPercent) * opts.EndMillDiameter
	outerLayer := opts.MaterialDiameter/2 + opts.EndMillDiameter/2
	offModel := outerLayer + 2*opts.EndMillDiameter
	if err := move(opts, gcode, 0, 0, 0); err != nil {
		return err
	}
	if err := gcode.SpindleOn(); err != nil {
		return err
	}
	if opts.FlatLayers == 0 {
		if err := move(opts, gcode, outerLayer, 0, math.NaN()); err != nil {
			return err
		}
	}
	for i := uint64(0); i < opts.FlatLayers; i++ {
		if err := millLayer(opts, gcode, layerWidth, outerLayer, offModel, -opts.LayerHeight*float64(i), 0); err != nil {
			return err
		}
	}
	z := -float64(opts.FlatLayers) * opts.LayerHeight
	for i := opts.StartZ; i <= opts.EndZ; i += opts.LayerHeight {
		noseRadius, err := opts.Profile.GetRadius(i, opts.Diameter, opts.Height)
		if err != nil {
			return err
		}
		if err := millLayer(opts, gcode, layerWidth, outerLayer, offModel, z, noseRadius); err != nil {
			return err
		}
		z -= opts.LayerHeight
	}
	return nil
}
