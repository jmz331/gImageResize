package gImageResize

import (
	"fmt"
	"testing"
)

func TestMultiSize(t *testing.T) {
	Thumbnails("tmp/1.png", [2]int{240, 240}, [2]int{360, 360})
}

func TestSingleSize(t *testing.T) {
	imgPath, _ := Thumbnail("tmp/2.png", 240, 240)
	fmt.Println(imgPath)
}
