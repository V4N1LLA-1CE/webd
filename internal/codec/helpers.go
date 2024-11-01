package codec

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/v4n1lla-1ce/webd/internal/pipeline"
)

func SaveToDisk(data pipeline.PipelineData) pipeline.PipelineData {
	buf, ok := data.Value.(bytes.Buffer)
	if !ok {
		fmt.Errorf("value is not a bytes.Buffer")
		return data
	}

	// create output filename with .png extension into the original given directory
	outputFilename := filepath.Join(data.Directory, data.BaseName+".png")

	err := os.WriteFile(outputFilename, buf.Bytes(), 0644)
	if err != nil {
		fmt.Errorf("failed to write file: %v\n", err)
		return data
	}

	data.Value = outputFilename
	return data
}
