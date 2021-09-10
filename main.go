package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/d4v1dw3bb/qtalkgoexample/child"
	"github.com/d4v1dw3bb/qtalkgoexample/parent"
	"github.com/progrium/qtalk-go/mux/frame"
)

func main() {
	flagDebug := flag.Bool("debug", false, "set's debug output")
	programType := flag.String("type", "parent", "Program Type (parent/sub)")
	flag.Parse()

	if *flagDebug {
		frame.Debug = os.Stderr
		fmt.Printf("type: %s", *programType)
	}

	switch *programType {
	case "child":
		child.Run()
	case "parent":
		parent.Run()
	default:
		fmt.Printf("Wrong value for flag type: %s", *programType)
	}

}
