package utils

import (
	"encoding/base64"
	"errors"
	"os"
	"strings"
)

// SaveVideoToFile parse data to golang image
func SaveVideoToFile(dirname string, filename string, data string) error {
	if data[5:10] != "video" {
		return errors.New("Data is not an video")
	}
	err := os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(dirname + filename)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	encodedString := data[strings.Index(data, ",")+1:]
	unbased, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return err
	}
	f.Write(unbased)
	return nil
}
