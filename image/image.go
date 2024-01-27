package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	height, width int
}

func (img Image) Bounds() image.Rectangle {

	return image.Rect(0, 0, img.height, img.width)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := uint8(x ^ y)

	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{height: 50, width: 100}
	pic.ShowImage(m)
}
