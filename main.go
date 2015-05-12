package main

import (
	"github.com/disintegration/imaging"
	"image"
	"image/png"
	"jin/test/graphics"
	"os"
)

func main() {
	test("1")
	test("2")
}

func test(name string) {
	originFilePath := name + ".png"
	originImage, _ := imaging.Open(originFilePath)

	img1 := imaging.Thumbnail(originImage, 240, 240, imaging.Lanczos)
	imaging.Save(img1, name+"_1.png")

	img2 := imaging.Fit(originImage, 240, 240, imaging.Lanczos)
	imaging.Save(img2, name+"_2.png")

	img3 := imaging.Resize(originImage, 240, 240, imaging.Lanczos)
	imaging.Save(img3, name+"_3.png")

	img4 := imaging.Fit(originImage, 240, 240, imaging.Box)
	imaging.Save(img4, name+"_4.png")

	img5 := imaging.Resize(originImage, 240, 0, imaging.Lanczos)
	imaging.Save(img5, name+"_5.png")

	process(name, 240, 240)
}

func process(name string, width, height int) {
	img, _ := LoadImage(name + ".png")
	dst := image.NewRGBA(image.Rect(0, 0, 400, 300))

	if err := graphics.Thumbnail(dst, img); err != nil {
		panic(err)
	}

	// white := color.RGBA{255, 255, 255, 255}
	// imgNew := image.NewRGBA(db)
	// draw.Draw(imgNew, bounds, &image.Uniform{white, image.ZP, draw.Src)
	// draw.Draw(imgNew, bounds, img, image.ZP, draw.Src)

	file, err := os.Create(name + "_6.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, dst)
}

func LoadImage(filename string) (image.Image, error) {
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
