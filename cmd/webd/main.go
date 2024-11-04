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
		pipeline.Webp2Png(args)
		break

	case "png2webp":
		pipeline.PNG2Webp(args)
		break

	case "jpg2png":
		pipeline.Jpg2Png(args)
		break

	case "png2jpg":
		pipeline.Png2Jpg(args)
		break

	case "jpg2webp":
		pipeline.Jpg2Webp(args)
		break

	case "webp2jpg":
		pipeline.Webp2Jpg(args)
		break

	default:
		fmt.Printf("unsupported mode: %s\n", args.Mode)
	}

}
