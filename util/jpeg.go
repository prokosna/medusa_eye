package util

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/pixiv/go-libjpeg/jpeg"
)

// For debugging
func WriteFile(id uint64, data []byte) error {
	file, err := os.Create("images/image_" + fmt.Sprint(id) + ".jpg")
	defer file.Close()
	if err != nil {
		return err
	}
	img, err := jpeg.Decode(bytes.NewReader(data), &jpeg.DecoderOptions{})
	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	err = jpeg.Encode(w, img, &jpeg.EncoderOptions{Quality: 90})
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}
