package logic

import (
	"os"

	"../gcode"
	"../options"
)

// DoSlice performs a slice of the nose cone and saves the output file
func DoSlice(opts options.Options) error {
	var writer gcode.Writer
	if opts.OutputFile == nil {
		writer = gcode.CreateWriter(os.Stdout)
	} else {
		file, err := os.Open(*opts.OutputFile)
		if err != nil {
			return err
		}
		defer file.Close()
		writer = gcode.CreateWriter(file)
	}
	if err := Slice(opts, &writer); err != nil {
		return err
	}
	ShowStats(opts, writer)
	return nil
}
