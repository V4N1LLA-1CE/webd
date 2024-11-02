package main

import (
	_ "embed"
	"fmt"

	"github.com/v4n1lla-1ce/webd/internal/cli"
	"github.com/v4n1lla-1ce/webd/internal/codec"
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
	ConvertWebpToPNG(args.DirPath, args.DeleteOrigin)
}

func ConvertWebpToPNG(directory string, deleteWebp bool) {
	fmt.Printf("Starting WebP to PNG conversion\n")
	fmt.Printf("Source directory: %v\n", directory)
	fmt.Printf("Mode: %s\n", map[bool]string{
		true:  "Converting WebP to PNG and deleting original files",
		false: "Converting WebP to PNG and preserving original files",
	}[deleteWebp])
	fmt.Println()

	// load data into pipeline
	files := pipeline.LoadPipeline(directory, deleteWebp)

	// decode webp to raw image format
	decoded := pipeline.NewPipeline(files, codec.DecodeWebp)

	// encode raw image as png
	encoded := pipeline.NewPipeline(decoded, codec.EncodeToPng)

	// save images to disk
	saved := pipeline.NewPipeline(encoded, codec.SaveToDisk)

	for result := range saved {
		if outPath, ok := result.Value.(string); ok {
			if deleteWebp {
				fmt.Printf("Converted WebP to PNG and saved to: %s\n - Deleted original: %v\n", outPath, result.SourcePath)
			} else {
				fmt.Printf("Converted WebP to PNG and saved to: %s\n", outPath)
			}
		}
	}
}
