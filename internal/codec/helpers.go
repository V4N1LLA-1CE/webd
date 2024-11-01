package codec

import (
	"bytes"
	"fmt"
	"os"
)

func SaveToDisk(img bytes.Buffer, path string) (string, error) {
	// create file path
	filepath := fmt.Sprintf("%v.png", path)

	err := os.WriteFile(filepath, img.Bytes(), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v\n", err)
	}
	return filepath, nil
}
