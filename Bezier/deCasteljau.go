package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

var (
	n      int
	x, y   []int
	img    = image.NewRGBA(image.Rect(0, 0, 640, 480))
	drawer = func(x, y int) {
		img.Set(x, y, color.RGBA{A: 255})
	}
)

func one_bezier(a, b int, t float64) float64 {
	return (1-t)*float64(a) + t*float64(b)
}

func bezier_curve(x []int, n, k int, t float64) float64 {
	if n == 1 {
		return one_bezier(x[k], x[k+1], t)
	} else {
		return (1-t)*bezier_curve(x, n-1, k, t) + t*bezier_curve(x, n-1, k+1, t)
	}
}

func bezier(x, y []int, num int, bx, by *[]int) {
	n := len(x) - 1
	step := 1.0 / float64(num-1)
	tStep := make([]float64, 5)
	args := 0.0
	for true {
		if args > 1+step {
			break
		}
		tStep = append(tStep, args)
		args += step
	}
	for _, t := range tStep {
		*bx = append(*bx, int(bezier_curve(x, n, 0, t)))
		*by = append(*by, int(bezier_curve(y, n, 0, t)))
	}
}

func main() {
	n = 6
	x = []int{0, 200, 500, 600, 250, 300}
	y = []int{0, 160, 180, 250, 400, 380}
	for i := 0; i < n; i++ {
		drawer(x[i], y[i])
	}

	var (
		b_x, b_y []int
	)

	bezier(x, y, 10000, &b_x, &b_y)

	for i := 0; i < len(b_x); i++ {
		drawer(b_x[i], b_y[i])
	}

	imgfile, _ := os.Create("Bezier.png")
	defer imgfile.Close()
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}
