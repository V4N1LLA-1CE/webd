package codec

import (
	"bytes"
	"fmt"
	"golang.org/x/image/webp"
	"image"
	"image/png"
	"os"
)

func DecodeWebp(path string) image.Image {
	// open file
	f, err := os.Open(path)
	if err != nil {
		fmt.Errorf("failed to open file: %v", err)
		return nil
	}
	defer f.Close()

	// decode webp
	img, err := webp.Decode(f)
	if err != nil {
		fmt.Errorf("failed to decode webp: %v", err)
		return nil
	}

	return img
}

func EncodeToPng(img image.Image) bytes.Buffer {
	// make a buffer
	var buf bytes.Buffer

	// encode and write to buffer
	err := png.Encode(&buf, img)
	if err != nil {
		fmt.Errorf("failed to encode to png: %v", err)
		return buf
	}

	return buf
}
