package main

import (
	"fmt"
	"github.com/v4n1lla-1ce/webd/internal/cli"
	"github.com/v4n1lla-1ce/webd/internal/codec"
	"github.com/v4n1lla-1ce/webd/internal/pipeline"
)

func main() {
	// get arguments from cli
	args := cli.GetArgs()

	fmt.Printf("Converting all webp in %v to png\n", args.DirPath)

	// load data into pipeline
	webpImages := pipeline.LoadPipeline(args.DirPath)

	// decode webp to raw image format
	rawImages := pipeline.NewPipeline(webpImages, codec.DecodeWebp)

	// encode raw image as png / jpeg
	pngImages := pipeline.NewPipeline(rawImages, codec.EncodeToPng)

	// save images to disk
	filenames := pipeline.NewPipeline(pngImages, codec.SaveToDisk)
	for name := range filenames {
		fmt.Println(name)
	}
}
