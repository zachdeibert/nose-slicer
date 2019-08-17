package options

// Options represents all of the possible options a user can specify from the
// command line
type Options struct {
	Mode               Mode
	Profile            Profile
	Diameter           float64
	Height             float64
	StartZ             float64
	EndZ               float64
	LayerHeight        float64
	FeedRate           float64
	RapidFeedRate      float64
	LayerDwell         float64
	MaterialDiameter   float64
	EndMillDiameter    float64
	Unit               Unit
	Direction          Direction
	PassOverlapPercent float64
	FlatLayers         uint64
	OutputFile         *string
}
