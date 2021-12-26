package manipulate

import (
	"errors"
	"fmt"
	"github.com/Festivals-App/festivals-fileserver/server/config"
	"github.com/disintegration/imaging"
	"image"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func ResizeIfNeeded(conf *config.Config, objectID string, parameters url.Values) (*os.File, error) {

	// get values
	widthVal := parameters.Get("width")
	heightVal := parameters.Get("height")

	// declare needed vars
	var err error
	var width = 0
	var height = 0

	if widthVal != "" {
		width, err = strconv.Atoi(widthVal)
		if err != nil {
			return nil, err
		}
	}

	if heightVal != "" {
		height, err = strconv.Atoi(heightVal)
		if err != nil {
			return nil, err
		}
	}

	if !validImageDimensions(width, height) {
		return nil, errors.New("the provided image dimensions are invalid")
	}

	// create path to resized image
	var prefix string
	if width > height {
		prefix = fmt.Sprintf("w%d_", width)
	} else {
		prefix = fmt.Sprintf("h%d_", height)
	}
	resizedImageName := prefix + objectID
	resizedImagePath := filepath.Join(conf.ResizeStorageURL, resizedImageName)

	// check if resizded image already exists
	if FileExists(resizedImagePath) {
		resizedImage, err := os.Open(resizedImagePath)
		if err != nil {
			return nil, err
		}
		return resizedImage, nil
	}

	return Resize(conf, objectID, width, height)
}

func Resize(conf *config.Config, objectID string, width int, height int) (*os.File, error) {

	// get original image
	originalImagePath := filepath.Join(conf.StorageURL, objectID)
	srcImage, err := imaging.Open(originalImagePath)
	if err != nil {
		return nil, err
	}
	// resize image
	resizedImage := imaging.Fit(srcImage, width, height, imaging.CatmullRom)
	// make resize dir if it does not exist
	_ = os.MkdirAll(conf.ResizeStorageURL, os.ModePerm)
	// create path to resized image
	var prefix string
	if width > height {
		prefix = fmt.Sprintf("w%d_", width)
	} else {
		prefix = fmt.Sprintf("h%d_", height)
	}
	resizedImageName := prefix + objectID
	resizedImagePath := filepath.Join(conf.ResizeStorageURL, resizedImageName)
	// save the file
	err = imaging.Save(resizedImage, resizedImagePath, imaging.JPEGQuality(75))
	if err != nil {
		return nil, err
	}
	// read and return file
	img, err := os.Open(resizedImagePath)
	if err != nil {
		return nil, err
	}
	return img, nil
}
