package pipeline

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type PipelineData struct {
	Value        any
	SourcePath   string
	Directory    string
	BaseName     string
	DeleteOrigin bool
}

func LoadPipeline(directory string, deleteOrigin bool) <-chan PipelineData {
	out := make(chan PipelineData)

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
				baseName := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))

				data := PipelineData{
					SourcePath:   fullpath,
					Directory:    directory,
					BaseName:     baseName,
					DeleteOrigin: deleteOrigin,
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

	// spawn goroutine in background to convert image
	go func() {
		for in := range input {
			out <- process(in)
		}
		close(out)
	}()

	return out
}
