package internal

import (
	"bytes"
	"embed"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	//"os"

	"github.com/mattn/go-sixel"
)

//go:embed images/*.jpeg
var embeddedImages embed.FS

// ProcessEmbeddedImage processes an embedded image and returns it as sixel data
func ProcessEmbeddedImage(imageName string) (string, error) {
	f, err := embeddedImages.Open("images/" + imageName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return processImage(f)
}

// func ProcessEmbeddedImage(imageName string) error {
// 	f, err := embeddedImages.Open("images/" + imageName)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	return processImage(f)
// }

// ProcessImageFile processes an external image file and returns it as sixel data
// func ProcessImageFile(imagePath string) (string, error) {
// 	f, err := embeddedImages.Open(imagePath)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer f.Close()

// 	return processImage(f)
// }

// processImage handles the common image processing logic
func processImage(r io.Reader) (string, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	enc := sixel.NewEncoder(&buf)
	err = enc.Encode(img)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// func processImage(r io.Reader) error {
// 	img, _, err := image.Decode(r)
// 	if err != nil {
// 		return err
// 	}

// 	enc := sixel.NewEncoder(os.Stdout)
// 	return enc.Encode(img)
// }

func ListEmbeddedImages() ([]string, error) {
	var images []string
	entries, err := embeddedImages.ReadDir("images")
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			images = append(images, entry.Name())
		}
	}
	return images, nil
}
