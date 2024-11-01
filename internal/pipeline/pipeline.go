package pipeline

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadPipeline(directory string) <-chan string {
	out := make(chan string)

	go func() {
		// close channel once there are no more webp files to be added to pipeline
		defer close(out)

		entries, err := os.ReadDir(directory)
		if err != nil {
			fmt.Errorf("Error reading directory: %v\n", err)
			return
		}

		// find all webp and add to channel
		for _, entry := range entries {
			// skip directories
			if entry.IsDir() {
				continue
			}

			// check if webp, if yes, add full path to channel
			if strings.ToLower(filepath.Ext(entry.Name())) == ".webp" {
				fullpath := filepath.Join(directory, entry.Name())
				out <- fullpath
			}
		}
	}()

	return out
}

func NewPipeline[I any, O any](input <-chan I, process func(I) O) <-chan O {
	// make channel
	out := make(chan O)

	// spawn goroutine in background to convert image
	go func() {
		for in := range input {
			out <- process(in)
		}
		close(out)
	}()

	return out
}
