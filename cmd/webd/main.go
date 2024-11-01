package main

import (
	"fmt"
	"os"
)

type Arguments struct {
	Dir string
}

func getArgs() Arguments {
	if len(os.Args) != 2 {
		fmt.Println("Usage: webd <directory>")
		os.Exit(1)
	}

	return Arguments{
		Dir: os.Args[1],
	}
}

func main() {
	// get arguments from cli
	args := getArgs()

	fmt.Printf("Converting all webp in %v to png\n", args.Dir)

	// scan directory
	// put job into pipeline
}
