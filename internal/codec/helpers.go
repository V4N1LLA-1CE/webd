package codec

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/v4n1lla-1ce/webd/internal/types"
)

func SaveToDisk(data types.PipelineData) types.PipelineData {
	buf, ok := data.Value.(bytes.Buffer)
	if !ok {
		fmt.Errorf("value is not a bytes.Buffer")
		return data
	}

	// create output filename with .png extension into the original given directory
	outputFilename := filepath.Join(data.Directory, data.BaseName+"."+data.TargetExt)

	err := os.WriteFile(outputFilename, buf.Bytes(), 0644)
	if err != nil {
		fmt.Errorf("failed to write file: %v\n", err)
		return data
	}

	// check flag
	if data.DeleteOrigin == true {
		// create file path
		err := os.Remove(data.SourcePath)
		if err != nil {
			fmt.Errorf("failed to delete original file: %s: %v\n", data.SourcePath, err)
		}
	}

	data.Value = outputFilename
	return data
}
