package codec

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/chai2010/webp"
	"github.com/v4n1lla-1ce/webd/internal/types"
)

func DecodeWebp(data types.PipelineData) types.PipelineData {
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

func EncodeToWebp(data types.PipelineData) types.PipelineData {
	// assert value in data to be image.Image
	img, ok := data.Value.(image.Image)
	if !ok {
		fmt.Errorf("value is not an image")
		return data
	}

	// make buffer
	var buf bytes.Buffer

	// encode img and write to buf
	err := webp.Encode(&buf, img, &webp.Options{
		Lossless: false,
		Quality:  95,
	})
	if err != nil {
		fmt.Errorf("failed to encode to webp: %v", err)
	}

	data.Value = buf
	return data
}

func DecodePng(data types.PipelineData) types.PipelineData {
	f, err := os.Open(data.SourcePath)
	if err != nil {
		fmt.Errorf("failed to open file: %v", err)
		return data
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		fmt.Errorf("failed to decode png: %v", err)
	}

	data.Value = img
	return data
}

func EncodeToPng(data types.PipelineData) types.PipelineData {
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
