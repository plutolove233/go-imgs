// Package main
/*
@Coding : utf-8
@time : 2022/8/29 20:38
@Author : yizhigopher
@Software : GoLand
*/
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type drawer func(x, y int)

func midPoint(x0, y0, x1, y1 int, brush drawer) {
	sx, sy := 1, 1
	a := y0 - y1
	b := x1 - x0
	d, d1, d2 := 2*a+b, 2*a, 2*(a+b)

	x, y := x0, y0
	for {
		if x == x1 {
			return
		}
		brush(x, y)
		if d < 0 {
			x += sx
			y += sy
			d += d2
		} else {
			x += sx
			d += d1
		}
	}
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	midPoint(0, 0, 500, 200, func(x, y int) {
		img.Set(x, y, color.RGBA{0, 0, 0, 255})
	})

	imgFile, _ := os.Create("midPoint.png")
	defer imgFile.Close()
	err := png.Encode(imgFile, img)
	if err != nil {
		log.Fatal(err)
	}
}
