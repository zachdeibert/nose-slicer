package options

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
