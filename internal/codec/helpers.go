package codec

import (
	"bytes"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func SaveToDisk(img bytes.Buffer) string {
	filename := fmt.Sprintf("%v.png", uuid.New().String())
	os.WriteFile(filename, img.Bytes(), 0644)

	return filename
}
