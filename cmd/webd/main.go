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
	files := pipeline.LoadPipeline(args.DirPath)

	// decode webp to raw image format
	decoded := pipeline.NewPipeline(files, codec.DecodeWebp)

	// encode raw image as png
	encoded := pipeline.NewPipeline(decoded, codec.EncodeToPng)

	// save images to disk
	saved := pipeline.NewPipeline(encoded, codec.SaveToDisk)

	for result := range saved {
		if outPath, ok := result.Value.(string); ok {
			fmt.Printf("Converted and saved to: %s\n", outPath)
		}
	}
}
