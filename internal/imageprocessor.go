package internal

import (
	"embed"
	"io"
	"io/ioutil"
)

//go:embed images/*.jpeg
var embeddedImages embed.FS

// ProcessEmbeddedImage processes an embedded ascii image and returns it as a string
// using the processImage function
func ProcessEmbeddedImage(imageName string) (string, error) {
	f, err := embeddedImages.Open("images/" + imageName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return processImage(f)
}

// processImage handles the common image processing logic
// uses io.Reader to read the image data that is in ascii format
// and returns it as a string
func processImage(r io.Reader) (string, error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

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
