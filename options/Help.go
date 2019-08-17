package options

const (
	// HelpMessage is the message to show for help with arguments
	HelpMessage string = `Usage: nose-slicer [options]

Options:

    -h, --help              Shows this help message
    -v, --version           Shows the version of this program
    -p, --profile           Sets the profile of the nose cone to slice
                            [conical|ogive]
    -nd, --nose-diameter    Sets the diameter of the nose cone
    -nh, --nose-height      Sets the height of the nose cone
    -z0, --start-z          Sets the position (from 0 at the tip to
                            {nose-height} at the bottom where the cut starts)
    -z1, --end-z            Sets the position (from 0 at the tip to
                            {nose-height} at the bottom where the cut starts)
    -l, --layer-height      Sets the height of each layer of the cut
    -F, --feed-rate         Sets the feed rate for cutting
    -G1, --rapid-feed-rate  Sets the feed rate for moving (if none is specified
                            rapid move commands are used, but then time
                            estimates are not accurate)
    -D, --dwell             Sets how long the CNC should wait after each layer
                            to allow the part and end mill to cool down
    -d, --mat-diameter      Sets the diameter of the material that will be cut.
    -e, --end-mill-diameter Sets the diameter of the end mill
    -u, --unit              Sets the unit that all parameters are in, and if the
                            emitted GCode will use metric or imperial units
                            [um|micron(s)|mm|millimeter(s)|cm|centimeter(s)|m|
                            meter(s)|thou|mil(s)|in|inch(es)]
    -um, -mm, -cm, -m, -in  Shortcuts for setting the unit (see -u, --unit)
    -cw, -ccw               Sets the direction in which the circles are milled
    -O, --overlay           Sets the percentage of each cut that will overlay
                            the last
    -f, --flat-layers       Sets the number of layers to cut completely flat at
                            the top to make sure the tip will be inside the
                            existing material
    -o, --out               Sets the file to output the GCode to.  If not set,
                            the GCode will be printed to standard output.

Required Options:

    --profile, --nose-diameter (or --mat-diameter), --nose-height,
    --end-mill-diameter

Default Values:

    --nose-diameter         The same as --mat-diameter
    --start-z               0
    --end-z                 The same as --nose-height
    --layer-height          0.5
    --feed-rate             1000
    --rapid-feed-rate       Use G0 commands
    --dwell                 None
    --mat-diameter          The same as --nose-diameter
    --unit                  Millimeters
    -cw                     Clockwise
    --overlay               30%
    --flat-layers           2
    --out                   Standard Output
`
)
