package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

var img = image.NewRGBA(image.Rect(0, 0, 640, 480))
var drawer = func(x, y int) {
	img.Set(x, y, color.RGBA{0, 0, 0, 255})
}

func ClipT(p, q float64, u1, u2 *float64) bool {
	var r float64
	if p < 0 {
		r = q / p
		if r > *u2 {
			return false
		}
		if r > *u1 {
			*u1 = r
		}
	} else if p > 0 {
		r = q / p
		if r < *u1 {
			return false
		}
		if r < *u2 {
			*u2 = r
		}
	} else {
		return q >= 0
	}
	return true
}

func ABS(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func drawLine(x0, y0, x1, y1 int) {
	sx := 1
	sy := 1
	dx := ABS(x1 - x0)
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	dy := ABS(y1 - y0)
	var e int
	if dx > dy {
		e = dx
	} else {
		e = -dy
	}
	for {
		drawer(x0, y0)
		if x0 == x1 && y0 == y1 {
			return
		}
		if e > -dx {
			e -= dy * 2
			x0 += sx
		}
		if e < dy {
			e += dx * 2
			y0 += sy
		}
	}
}

func LB_LineClip(x1, y1, x2, y2, XL, XR, YB, YT float64) {
	dx := x2 - x1
	dy := y2 - y1
	var t0, t1 float64
	t0 = 0
	t1 = 1
	if ClipT(-dx, x1-XL, &t0, &t1) {
		if ClipT(dx, XR-x1, &t0, &t1) {
			if ClipT(-dy, y1-YB, &t0, &t1) {
				if ClipT(dy, YT-y1, &t0, &t1) {
					drawLine(int(x1+t0*dx), int(y1+t0*dy),
						int(x1+t1*dx), int(y1+t1*dy))
				}
			}
		}
	}
}

func main() {
	XL := 100
	XR := 500
	YB := 100
	YT := 380
	drawLine(XL, YB, XL, YT)
	drawLine(XL, YB, XR, YB)
	drawLine(XL, YT, XR, YT)
	drawLine(XR, YB, XR, YT)
	LB_LineClip(50, 110, 200, 300, float64(XL), float64(XR), float64(YB), float64(YT))
	imgfile, _ := os.Create("LiangBarsky.png")
	defer imgfile.Close()
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}
