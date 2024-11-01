package main

import (
	"fmt"
	"github.com/v4n1lla-1ce/webd/internal/cli"
)

func main() {
	// get arguments from cli
	args := cli.GetArgs()

	fmt.Printf("Converting all webp in %v to png\n", args.Dir)

	// put job into pipeline
}
