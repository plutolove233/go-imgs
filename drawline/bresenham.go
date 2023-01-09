package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	Bresenham(0, 0, 500, 400, func(x, y int) {
		img.Set(x, y, color.RGBA{0, 0, 0, 255})
	})
	imgfile, _ := os.Create("Bresenham.png")
	defer imgfile.Close()
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

func Bresenham(x0, y0, x1, y1 int, f func(x int, y int)) {
	dx := x1 - x0
	dy := y1 - y0
	e := -dx
	x, y := x0, y0
	for i := 0; i <= dx; i++ {
		f(x, y)
		x++
		e += 2 * dy
		if e >= 0 {
			y++
			e -= 2 * dx
		}
	}
}
