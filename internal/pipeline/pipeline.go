package pipeline

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/v4n1lla-1ce/webd/internal/cli"
	"github.com/v4n1lla-1ce/webd/internal/codec"
	"github.com/v4n1lla-1ce/webd/internal/types"
)

func LoadPipeline(args cli.Arguments) <-chan types.PipelineData {
	// directory string, deleteOrigin bool
	out := make(chan types.PipelineData)

	go func() {
		// close channel once there are no more webp files to be added to pipeline
		defer close(out)

		entries, err := os.ReadDir(args.DirPath)
		if err != nil {
			fmt.Errorf("Error reading directory: %v\n", err)
			return
		}

		// skip directories
		for _, entry := range entries {
			// skip directories
			if entry.IsDir() {
				continue
			}

			// check if webp, if yes, add full path to channel
			if strings.ToLower(filepath.Ext(entry.Name())) == args.Ext {
				fullpath := filepath.Join(args.DirPath, entry.Name())
				baseName := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))

				data := types.PipelineData{
					SourcePath:   fullpath,
					Directory:    args.DirPath,
					BaseName:     baseName,
					DeleteOrigin: args.DeleteOrigin,
				}

				out <- data
			}
		}
	}()

	return out
}

func NewPipeline[I any, O any](input <-chan I, process func(I) O) <-chan O {
	// make channel
	out := make(chan O)

	// spawn goroutine in background
	go func() {
		for in := range input {
			out <- process(in)
		}
		close(out)
	}()

	return out
}

func ConvertWebpToPNG(args cli.Arguments) {

	fmt.Printf("Starting WebP to PNG conversion\n")
	fmt.Printf("Source directory: %v\n", args.DirPath)
	fmt.Printf("Mode: %s\n", map[bool]string{
		true:  "Converting WebP to PNG and deleting original files",
		false: "Converting WebP to PNG and preserving original files",
	}[args.DeleteOrigin])
	fmt.Println()

	// make channel to check to see if any files are processed
	done := make(chan struct{})
	var filesFound bool
	var count int

	// time benchmarks
	var start time.Time
	var elapsed time.Duration

	// background goroutine to count number of files processed
	go func() {

		// start timing before loading data
		start = time.Now()

		// load data into pipeline
		files := LoadPipeline(args)

		// decode webp to raw image format
		decoded := NewPipeline(files, codec.DecodeWebp)

		// encode raw image as png
		encoded := NewPipeline(decoded, codec.EncodeToPng)

		// save images to disk
		saved := NewPipeline(encoded, codec.SaveToDisk)

		for result := range saved {
			filesFound = true
			if outPath, ok := result.Value.(string); ok {
				count++
				if args.DeleteOrigin && args.Verbosity {
					fmt.Printf("Converted WebP to PNG and saved to: %s\n - Deleted original: %v\n", outPath, result.SourcePath)
				} else if args.Verbosity {
					fmt.Printf("Converted WebP to PNG and saved to: %s\n", outPath)
				}
			}
		}
		// get time elapsed since beginning
		elapsed = time.Since(start)

		// use empty struct as a signal since it takes 0 bytes of mem
		done <- struct{}{}
	}()
	// wait on main thread for signal
	<-done

	if !filesFound {
		fmt.Println("No images have been processed in the specified directory")
	} else {
		fmt.Printf("%v files converted\n", count)
		fmt.Printf("Processing time: %.3f seconds\n", elapsed.Seconds())
	}
}
