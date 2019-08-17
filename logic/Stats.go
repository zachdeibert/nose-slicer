package logic

import (
	"fmt"
	"io"
	"math"
	"os"

	"../gcode"
	"../options"
)

// ShowStats prints stats about the model to the console
func ShowStats(opts options.Options, gcode gcode.Writer) {
	var stream io.Writer
	if opts.OutputFile == nil {
		stream = os.Stderr
	} else {
		stream = os.Stdout
	}
	fmt.Fprintf(stream, `Nose cone stats:
Milling time: %.2f
Number of layers: %d
Starting material: %.1f-radius, %.1f-tall (at least) cylinder
Start with machine zeroed at the center of the top of the cylinder
`, gcode.Time, int(math.Ceil(opts.Height/opts.LayerHeight)), opts.MaterialDiameter/2, opts.Height)
}
