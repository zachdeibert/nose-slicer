package options

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Parse parses all of the options from the command line
func (opts *Options) Parse(args []string) error {
	var err error
	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			opts.Mode = PrintHelp
		case "-v", "--version":
			opts.Mode = PrintVersion
		case "-p", "--profile":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			switch args[i] {
			case "conical":
				opts.Profile = Conical
			case "ogive":
				opts.Profile = Ogive
			default:
				return errors.New("Invalid profile type")
			}
		case "-nd", "--nosediameter", "--nose-diameter":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.Diameter, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid diameter: %s", err.Error())
			}
			if opts.Diameter <= 0 {
				return errors.New("Diameter must be positive")
			}
		case "-nh", "--noseheight", "--nose-height":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.Height, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid height: %s", err.Error())
			}
			if opts.Height <= 0 {
				return errors.New("Height must be positive")
			}
		case "-z0", "--startz", "--start-z":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.StartZ, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid starting Z: %s", err.Error())
			}
			if opts.StartZ <= 0 {
				return errors.New("Starting Z must be positive")
			}
		case "-z1", "--endz", "--end-z":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.EndZ, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid ending Z: %s", err.Error())
			}
			if opts.EndZ <= 0 {
				return errors.New("Ending Z must be positive")
			}
		case "-l", "--layerheight", "--layer-height":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.LayerHeight, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid layer height: %s", err.Error())
			}
			if opts.LayerHeight <= 0 {
				return errors.New("Layer height must be positive")
			}
		case "-F", "--feedrate", "--feed-rate":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.FeedRate, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid feed rate: %s", err.Error())
			}
			if opts.FeedRate <= 0 {
				return errors.New("Feed rate must be positive")
			}
		case "-G1", "--rapidfeedrate", "--rapid-feed-rate":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.RapidFeedRate, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid rapid feed rate: %s", err.Error())
			}
			if opts.RapidFeedRate <= 0 {
				return errors.New("Rapid feed rate must be positive")
			}
		case "-D", "--dwell":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.LayerDwell, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid dwell: %s", err.Error())
			}
			if opts.LayerDwell <= 0 {
				return errors.New("Dwell must be positive")
			}
		case "-d", "--matdiameter", "--mat-diameter":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.MaterialDiameter, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid material diameter: %s", err.Error())
			}
			if opts.MaterialDiameter <= 0 {
				return errors.New("Material diameter must be positive")
			}
		case "-e", "--endmilldiameter", "--end-mill-diameter":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.EndMillDiameter, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid end mill diameter: %s", err.Error())
			}
			if opts.EndMillDiameter <= 0 {
				return errors.New("End mill diameter must be positive")
			}
		case "-u", "--unit":
			i++
			switch args[i] {
			case "um", "micron", "microns":
				opts.Unit = Microns
			case "mm", "millimeter", "millimeters":
				opts.Unit = Millimeters
			case "cm", "centimeter", "centimeters":
				opts.Unit = Centimeters
			case "m", "meter", "meters":
				opts.Unit = Meters
			case "thou", "mil", "mils":
				opts.Unit = Thou
			case "in", "inch", "inches":
				opts.Unit = Inches
			default:
				return errors.New("Invalid unit")
			}
		case "-um":
			opts.Unit = Microns
		case "-mm":
			opts.Unit = Millimeters
		case "-cm":
			opts.Unit = Centimeters
		case "-m":
			opts.Unit = Meters
		case "-in":
			opts.Unit = Inches
		case "-cw":
			opts.Direction = Clockwise
		case "-ccw":
			opts.Direction = Counterclockwise
		case "-O", "--overlap":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if strings.HasSuffix(args[i], "%") && !strings.HasSuffix(args[i], "%%") {
				args[i] = strings.TrimRight(args[i], "%")
			}
			if opts.PassOverlapPercent, err = strconv.ParseFloat(args[i], 64); err != nil {
				return fmt.Errorf("Invalid pass overlap percent: %s", err.Error())
			}
			opts.PassOverlapPercent /= 100
			if opts.PassOverlapPercent < 0 || opts.PassOverlapPercent > 1 {
				return errors.New("Pass overlap must be a percent")
			}
		case "-f", "--flatlayers", "--flat-layers":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			if opts.FlatLayers, err = strconv.ParseUint(args[i], 10, 64); err != nil {
				return fmt.Errorf("Invalid number of flat layers: %s", err.Error())
			}
			if opts.FlatLayers < 0 {
				return errors.New("Number of flat layers must not be negative")
			}
		case "-o", "--out":
			i++
			if i == len(args) {
				return fmt.Errorf("Expected parameter to %s", args[i-1])
			}
			opts.OutputFile = &args[i]
		default:
			return fmt.Errorf("Unknown argument '%s'", args[i])
		}
	}
	if opts.Profile == 0 {
		return fmt.Errorf("Must specify a profile (-p, --profile)")
	}
	if math.IsNaN(opts.Diameter) {
		if math.IsNaN(opts.MaterialDiameter) {
			return fmt.Errorf("Must specify a nose cone diameter (-nd, --nose-diameter)")
		} else {
			opts.Diameter = opts.MaterialDiameter
		}
	}
	if math.IsNaN(opts.Height) {
		return fmt.Errorf("Must specify a nose cone height (-nh, --nose-height)")
	}
	if math.IsNaN(opts.EndZ) {
		opts.EndZ = opts.Height
	}
	if math.IsNaN(opts.MaterialDiameter) {
		opts.MaterialDiameter = opts.Diameter
	}
	if math.IsNaN(opts.EndMillDiameter) {
		return fmt.Errorf("Must specify an end mill diameter (-e, --end-mill-diameter)")
	}
	return nil
}
