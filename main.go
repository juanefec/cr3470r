package main

import (
	"encoding/hex"
	"image"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	outputWidth  = 960
	outputHeight = 540
)

func main() {
	img := createWierdStuff(outputWidth, outputHeight)

	f, _ := os.Create("image3mini.png")
	defer f.Close()
	png.Encode(f, img)
}

func createStuff(w, h int) image.Image {
	img, err := doStuff(w, h)
	if err != nil {
		panic(err)
	}
	return img
}

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandStringBytesMaskImprSrc returns a random hexadecimal string of length n.
func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, (n+1)/2) // can be simplified to n/2 if n is always even

	if _, err := src.Read(b); err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)[:n]
}

func createWierdStuff(w, h int) image.Image {

	pixm := createPixelMatrix(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			rand := RandStringBytesMaskImprSrc(6)
			r, _ := strconv.ParseInt(rand[:2], 16, 64)
			g, _ := strconv.ParseInt(rand[2:4], 16, 64)
			b, _ := strconv.ParseInt(rand[4:6], 16, 64)

			//fmt.Println(rand[:2], rand[2:4], rand[4:6])
			pixm[x][y] = Pixel{
				int(r),
				int(g),
				int(b),
				255,
			}
		}
	}

	return imageFromPixels(pixm)

}

func createPixelMatrix(w, h int) [][]Pixel {
	pixs := make([][]Pixel, w)
	for x := 0; x < w; x++ {
		pixs[x] = make([]Pixel, h)
	}
	return pixs
}

func imageFromPixels(pixs [][]Pixel) image.Image {
	width := len(pixs)
	height := len(pixs[0])
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for x := 0; x < len(pixs); x++ {
		for y := 0; y < len(pixs[x]); y++ {
			img.Set(x, y, pixs[x][y])
		}
	}

	return img
}

// Get the bi-dimensional pixel array
func doStuff(w, h int) (image.Image, error) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {

		for x := 0; x < width; x++ {
			c := Pixel{
				int(Map(float64(y), 0, 2000, 200, 0)),
				int(Map(float64(x), 0, 2000, 0, 255)),
				int(Map(float64(x+y), 0, 5000, 0, 255)),
				255,
			}
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
