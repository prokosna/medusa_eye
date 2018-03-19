package util

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"os"
)

// For debugging
func WriteFile(id uint64, data []byte) error {
	file, err := os.Create("images/image_" + fmt.Sprint(id) + ".jpg")
	defer file.Close()
	if err != nil {
		return err
	}
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}
	return nil
}
