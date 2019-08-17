package options

// Direction represents the direction the cut should be made
type Direction int

const (
	// Clockwise is the way the clock goes
	Clockwise Direction = 1
	// Counterclockwise isn't the way the clock goes
	Counterclockwise Direction = 2
)
