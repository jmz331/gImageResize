package gImageResize

import (
	"errors"
	"fmt"
	"github.com/jmz331/gImageResize/graphics"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

func Thumbnails(src string, sizes ...[2]int) error {
	dist := src
	if len(sizes) > 0 {
		for _, size := range sizes {
			if err := Thumbnail(src, dist, size[0], size[1]); err != nil {
				return err
			}
		}
		return nil
	} else {
		return errors.New("缩略图尺寸不正确")
	}
}

func Thumbnail(src, dist string, width, height int) error {
	if width <= 0 || height <= 0 {
		return errors.New("缩略图尺寸不正确，应为大于0的正整数")
	}
	img, _ := loadImage(src)
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	if err := graphics.Thumbnail(dst, img); err != nil {
		return err
	}

	ext := path.Ext(dist)
	if dist == src {
		dist = fmt.Sprintf("%v_%v_%v%v", strings.TrimSuffix(dist, ext), width, height, ext)
		ext = strings.ToLower(ext)
	}
	file, err := os.Create(dist)
	if err != nil {
		return err
	}
	defer file.Close()
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
