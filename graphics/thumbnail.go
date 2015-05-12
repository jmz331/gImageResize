// Copyright 2011 The Graphics-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphics

import (
	"image"
	"image/color"
	"image/draw"
)

// Thumbnail scales and crops src so it fits in dst.
func Thumbnail(dst draw.Image, src image.Image) error {
	// Scale down src in the dimension that is closer to dst.
	sb := src.Bounds()
	db := dst.Bounds()
	rs := float64(sb.Dx()) / float64(sb.Dy())
	var b image.Rectangle

	bx := sb.Dx()
	by := sb.Dy()
	if bx > db.Dx() {
		bx = db.Dx()
		by = int(float64(bx) / rs)
	}
	if by > db.Dy() {
		by = db.Dy()
		bx = int(float64(by) * rs)
	}
	b = image.Rect(0, 0, bx, by)

	buf := image.NewRGBA(b)
	if err := Scale(buf, src); err != nil {
		return err
	}

	var pt image.Point
	if b.Dx() < db.Dx() {
		pt.X = (b.Dx() - db.Dx()) / 2
	} else {
		pt.Y = (b.Dy() - db.Dy()) / 2
	}
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(dst, db, &image.Uniform{white}, image.ZP, draw.Src)
	draw.Draw(dst, db, buf, pt, draw.Src)
	return nil
}
