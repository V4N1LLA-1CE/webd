package main

import (
	_ "embed"
	"fmt"

	"github.com/v4n1lla-1ce/webd/internal/cli"
	"github.com/v4n1lla-1ce/webd/internal/pipeline"
)

//go:embed banner.txt
var banner string

func main() {
	// print banner
	fmt.Println(banner)
	fmt.Printf("\n© Austin Sofaer (v4n1lla-1ce)\n\n")

	// get arguments from cli
	args := cli.GetArgs()

	// handle different conversion options
	switch args.Mode {
	case "webp2png":
		pipeline.Webp2PNG(args)
		break

	case "png2webp":
		fmt.Println("currently being implemented")
		break

	default:
		fmt.Printf("unsupported mode: %s\n", args.Mode)
	}

}
