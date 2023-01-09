package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Putpixel func(x, y int)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func DDA(x0, y0, x1, y1 int, brush Putpixel) {
	dx := abs(x0 - x1)
	dy := abs(y0 - y1)
	sx := 1
	if x0 >= x1 {
		sx = -1
	}
	var k float32 = (float32(dy) / float32(dx))
	var y float32 = float32(y0)
	for x := x0; x <= x1; x += sx {
		brush(x, int(y+0.5))
		y = y + k
	}
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	DDA(0, 0, 500, 200, func(x, y int) {
		img.Set(x, y, color.RGBA{0, 0, 0, 255})
	})
	imgfile, _ := os.Create("DDA.png")
	defer imgfile.Close()
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}
