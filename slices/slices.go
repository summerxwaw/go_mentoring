package main

import (
	"bytes"
	"encoding/base64"

	"image/png"

	"os"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	data := make([][]uint8, dy)
	for i := range data {
		data[i] = make([]uint8, dx)
	}

	for x := 0; x < dy; x++ {
		for y := 0; y < dx; y++ {
			data[x][y] = uint8(x * y)
		}
	}

	return data
}

func GenerateImage(img string) {
	b := img
	unbased, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		panic("Cannot decode b64")
	}

	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		panic("Bad png")
	}

	f, err := os.OpenFile("example.png", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Cannot open file")
	}

	png.Encode(f, im)
}

func main() {
	pic.Show(Pic)

	b := "iVBORw0KGgoAAAANSUhEUgAAADIAAABkCAIAAABLivQMAAAAwUlEQVR42uzaQQqDQAxG4US8/508ikdw6UKmlFJoFy6yKI3tB0FERILzePyEzBEjM7rVlBljRLfrfG9tiuN4ttnjvvfferS57+8H/L3n2Po1tl5r286l8vn3scVbvIWturfOal3r4an+fWzxFm9hS97CFm9h63/zVrWWxXwLW7yFLXkLW7zFW9gy38IWb2FL3pK3sMVb2LrufMv+FrZ4C1vyFrZ4i7ewZb6FLd7Clrwlb2GLt7BlfwtbF2TrFgAA//+xwOc680FROgAAAABJRU5ErkJggg=="

	unbased, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		panic("Cannot decode b64")
	}
	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		panic("Bad png")
	}
	f, err := os.OpenFile("example.png", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Cannot open file")
	}
	png.Encode(f, im)

}
