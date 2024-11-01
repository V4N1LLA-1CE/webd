package codec

import (
	"bytes"
	"fmt"
	"golang.org/x/image/webp"
	"image"
	"image/png"
	"os"
)

func DecodeWebp(path string) (image.Image, error) {
	// open file
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	// decode webp
	img, err := webp.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("failed to decode webp: %v", err)
	}

	return img, nil
}

func EncodeToPng(img image.Image) (bytes.Buffer, error) {
	// make a buffer
	var buf bytes.Buffer

	// encode and write to buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return buf, fmt.Errorf("failed to encode to png: %v", err)
	}

	return buf, nil
}
