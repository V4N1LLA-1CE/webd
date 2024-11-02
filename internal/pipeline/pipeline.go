package pipeline

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/v4n1lla-1ce/webd/internal/cli"
	"github.com/v4n1lla-1ce/webd/internal/types"
)

type ConversionPipeline struct {
	sourceExt string
	targetExt string
	decoder   func(types.PipelineData) types.PipelineData
	encoder   func(types.PipelineData) types.PipelineData
	saver     func(types.PipelineData) types.PipelineData
}

func LoadPipeline(args cli.Arguments) <-chan types.PipelineData {
	// build extension to search for
	ext := "." + strings.ToLower(args.SourceExt)

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

			if strings.ToLower(filepath.Ext(entry.Name())) == (ext) {
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

// handles any type of image conversion based on pipeline configuration
func (p *ConversionPipeline) Convert(args cli.Arguments) {
	fmt.Printf("Starting %s to %s conversion\n", strings.ToUpper(p.sourceExt), strings.ToUpper(p.targetExt))
	fmt.Printf("Source directory: %s\n", args.DirPath)
	fmt.Printf("Mode: %s\n", map[bool]string{
		true:  fmt.Sprintf("Converting %s to %s and deleting original files", strings.ToUpper(p.sourceExt), strings.ToUpper(p.targetExt)),
		false: fmt.Sprintf("Converting %s to %s and preserving original files", strings.ToUpper(p.sourceExt), strings.ToUpper(p.targetExt)),
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

		// decode to image
		decoded := NewPipeline(files, p.decoder)

		// encode image
		encoded := NewPipeline(decoded, p.encoder)

		// save images to disk
		saved := NewPipeline(encoded, p.saver)

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
