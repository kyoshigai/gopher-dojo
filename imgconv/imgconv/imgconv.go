package imgconv

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var ErrInvalidExt = errors.New("Cannot identify the extension")

type destExt string

type File struct {
	Path    string
	DestExt destExt
}

func isJPEG(path string) bool {
	return filepath.Ext(path) == ".jpeg" || filepath.Ext(path) == ".jpg"
}

func isPNG(path string) bool {
	return filepath.Ext(path) == ".png"
}

func Ext(path string) (destExt, error) {
	if isJPEG(path) {
		return destExt(".jpeg"), nil
	} else if isPNG(path) {
		return destExt(".png"), nil
	}
	return destExt(""), ErrInvalidExt
}

func (f File) Encode(destPath string, img image.Image) error {
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if isJPEG(destPath) {
		err = png.Encode(destFile, img)
		if err != nil {
			return err
		}
	} else if isPNG(destPath) {
		err = jpeg.Encode(destFile, img, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f File) Convert() error {
	srcFile, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	var img image.Image

	if isJPEG(f.Path) {
		img, err = jpeg.Decode(srcFile)
		if err != nil {
			return err
		}
	} else if isPNG(f.Path) {
		img, err = png.Decode(srcFile)
		if err != nil {
			return err
		}
	}

	_, srcFileName := filepath.Split(f.Path)
	baseName := strings.Replace(srcFileName, filepath.Ext(f.Path), "", 1)
	newFileName := "converted-" + baseName
	newFilePath := filepath.Dir(f.Path) + "/" + newFileName + string(f.DestExt)

	f.Encode(newFilePath, img)
	return nil
}
