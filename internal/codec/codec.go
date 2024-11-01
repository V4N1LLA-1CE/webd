package codec

import (
	"bytes"
	"fmt"
	"github.com/v4n1lla-1ce/webd/internal/pipeline"
	"golang.org/x/image/webp"
	"image"
	"image/png"
	"os"
)

func DecodeWebp(data pipeline.PipelineData) pipeline.PipelineData {
	// open file
	f, err := os.Open(data.SourcePath)
	if err != nil {
		fmt.Errorf("failed to open file: %v", err)
		return data
	}
	defer f.Close()

	// decode webp
	img, err := webp.Decode(f)
	if err != nil {
		fmt.Errorf("failed to decode webp: %v", err)
		return data
	}

	data.Value = img
	return data
}

func EncodeToPng(data pipeline.PipelineData) pipeline.PipelineData {
	// assert data.Value to be image.Image so encode works
	img, ok := data.Value.(image.Image)
	if !ok {
		fmt.Errorf("value is not an image")
		return data
	}

	// make a buffer
	var buf bytes.Buffer

	// encode and write to buffer
	err := png.Encode(&buf, img)
	if err != nil {
		fmt.Errorf("failed to encode to png: %v", err)
		return data
	}

	data.Value = buf
	return data
}
