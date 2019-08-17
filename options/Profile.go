package options

// Profile represents a cross-section type for the nose cone
type Profile int

const (
	// Conical is a nose cone that looks like a cone
	Conical Profile = 1
	// Ogive is a nose cone constructed from a circle
	Ogive Profile = 2
)
