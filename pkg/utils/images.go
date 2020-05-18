package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// DataToImage parse data to golang image
func DataToImage(data string) (image.Image, error) {
	if data[5:10] != "image" {
		return nil, errors.New("Data is not an image")
	}
	encodedString := data[strings.Index(data, ",")+1:]
	unbased, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(unbased)
	typeof := data[5:strings.Index(data, ";base64")]
	switch typeof {
	case "image/png":
		return png.Decode(reader)
	case "image/jpeg":
		return jpeg.Decode(reader)
	case "image/gif":
		return gif.Decode(reader)
	}
	return nil, errors.New("Image format not valid")
}

// SaveImageToFile parse data to golang image
func SaveImageToFile(dirname string, filename string, image image.Image, typeof string) error {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(dirname + filename)
	if err != nil {
		return err
	}
	switch typeof {
	case "image/png":
		return png.Encode(f, image)
	case "image/jpeg":
		opt := jpeg.Options{
			Quality: 90,
		}
		return jpeg.Encode(f, image, &opt)
	case "image/gif":
		opt := gif.Options{
			NumColors: 256,
		}
		return gif.Encode(f, image, &opt)
	}
	return errors.New("Image format not valid")
}
