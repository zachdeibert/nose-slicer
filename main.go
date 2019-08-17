package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"./logic"
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
		if err := logic.DoSlice(opts); err != nil {
			io.WriteString(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		panic(errors.New("Invalid mode"))
	}
}
