package main

import (
	"flag"
	"fmt"
)

func main() {
	flagDebug := flag.Bool("debug", false, "set's debug output")
	programType := flag.String("type", "parent", "Program Type (parent/sub)")
	flag.Parse()

	if *flagDebug {
		fmt.Println("test")
	}

	switch *programType {
	case "child":
		subprogram()
	case "parent":
		parentprogram()
	default:
		fmt.Printf("Wrong value for flag type: %s", *programType)
	}

}
