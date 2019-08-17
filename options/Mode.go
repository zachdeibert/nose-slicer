package options

// Mode represents something the program can do
type Mode int

const (
	// Normal slices a model
	Normal Mode = 0
	// PrintHelp shows the help and exits
	PrintHelp Mode = 1
	// PrintVersion shows the version and exits
	PrintVersion Mode = 2
)
