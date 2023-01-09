package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func circlePoint(x, y, p, q int, f func(x, y int)) {
	f(x+p, y+q)
	f(x-p, y+q)
	f(x+p, y-q)
	f(x-p, y-q)
	f(x+q, y+p)
	f(x-q, y+p)
	f(x+q, y-p)
	f(x-q, y-p)
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	midPointCircle(100, 100, 100, func(x, y int) {
		img.Set(x, y, color.RGBA{0, 0, 0, 255})
	})
	imgfile, _ := os.Create("circle.png")
	defer imgfile.Close()
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

func midPointCircle(xc, yc, r int, f func(x int, y int)) {
	x, y := 0, r
	d := 3 - 2*r
	circlePoint(xc, yc, x, y, f)
	for {
		if x >= y {
			break
		}
		if d < 0 {
			d += 4*x + 6
		} else {
			d += 4*(x-y) + 10
			y--
		}
		x++
		circlePoint(xc, yc, x, y, f)
	}
}
