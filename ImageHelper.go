package gImageResize

import (
	"errors"
	"github.com/jmz331/gImageResize/graphics"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

func Thumbnail(src, dist string, width, height int) error {
	if width <= 0 || height <= 0 {
		return errors.New("缩略图尺寸不正确，应为大于0的正整数")
	}
	img, _ := loadImage(src)
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	if err := graphics.Thumbnail(dst, img); err != nil {
		return err
	}

	file, err := os.Create(dist)
	if err != nil {
		return err
	}
	defer file.Close()
	ext := strings.ToLower(path.Ext(dist))
	switch ext {
	case "jpg", "jpeg":
		err = jpeg.Encode(file, dst, &jpeg.Options{Quality: 95})
	case "gif":
		err = gif.Encode(file, dst, &gif.Options{NumColors: 256})
	default:
		err = png.Encode(file, dst)
	}
	return err
}

func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
