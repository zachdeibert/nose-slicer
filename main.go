package main

import (
	"errors"
	"fmt"
	"os"

	"./options"
)

func main() {
	opts := options.DefaultOptions()
	if err := opts.Parse(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n\n%s", err.Error(), options.HelpMessage)
		os.Exit(1)
	}
	switch opts.Mode {
	case options.PrintHelp:
		fmt.Print(options.HelpMessage)
	case options.PrintVersion:
		fmt.Print(options.VersionString)
	case options.Normal:
		fmt.Println("// TODO")
	default:
		panic(errors.New("Invalid mode"))
	}
}
