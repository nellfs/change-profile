package main

import (
	"github.com/disintegration/imaging"
	"image"
)

func parseImage(imgPath string) (image.Image, error) {

	image, err := imaging.Open(imgPath)
	if err != nil {
		return nil, err
	}

	return image, err
}
