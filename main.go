package main

import (
	//"github.com/faiface/pixel"
	"image/png"
	//"github.com/faiface/pixel/pixelgl"
	"image"
	"os"
)

const (
	outputWidth  = 1920
	outputHeight = 1080
)

func main() {
	create(outputWidth, outputHeight)
}

func create(w, h int) {

	//pixs, err := getPixels(img)

	// do stuff

	img, err := doStuff(w, h)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("image2.png")
	png.Encode(f, img)
	defer f.Close()

}

// Get the bi-dimensional pixel array
func doStuff(w, h int) (image.Image, error) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {

		for x := 0; x < width; x++ {
			c := Pixel{int(Map(float64(y), 0, 2000, 200, 0)), int(Map(float64(x), 0, 2000, 0, 255)), int(Map(float64(x+y), 0, 5000, 0, 255)), 10}
			img.Set(x, y, c)
		}
	}

	return img, nil
}

// Get the bi-dimensional pixel array
func getPixels(img image.Image) ([][]Pixel, error) {

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func (p Pixel) RGBA() (uint32, uint32, uint32, uint32) {
	return uint32(p.R * 257), uint32(p.G * 257), uint32(p.B * 257), uint32(p.A * 257)
}

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}

func Map(v, s1, st1, s2, st2 float64) float64 {
	newval := (v-s1)/(st1-s1)*(st2-s2) + s2
	if s2 < st2 {
		if newval < s2 {
			return s2
		}
		if newval > st2 {
			return st2
		}
	} else {
		if newval > s2 {
			return s2
		}
		if newval < st2 {
			return st2
		}
	}
	return newval
}
