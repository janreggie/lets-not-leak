package main

import "errors"

// Image represents an instance of a "Let's not leak" image.
type Image struct {
	options Options
	width   int
	height  int
}

func NewImage(options Options) (*Image, error) {
	image := new(Image)
	image.width = 1920
	image.height = 1200
	image.options = options

	return image, nil
}

// Draw draws the image and saves it in memory. It must only be called once.
func (image *Image) Draw() error {
	return errors.New("unimplemented")
}

// Save saves the image to disk
func (image *Image) Save() error {
	return errors.New("unimplemented")
}
