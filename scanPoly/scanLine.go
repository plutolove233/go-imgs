package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type XET struct {
	ymax int
	x    float64
	dx   float64
	next *XET
}

type Point struct {
	x, y int
}

func PolyScan(polyPoint []Point, pointNum int, paint func(x, y int)) {
	maxY := 0
	for i := 0; i < pointNum; i++ {
		if polyPoint[i].y > maxY {
			maxY = polyPoint[i].y
		}
	}
	var (
		pAET = new(XET)
		pNET = make([]*XET, 1024)
	)
	pAET.next = nil
	for i := 0; i <= maxY; i++ {
		pNET[i] = new(XET)
		pNET[i].next = nil
	}

	for i := 0; i <= maxY; i++ {
		for j := 0; j < pointNum; j++ {
			if polyPoint[j].y == i {
				if polyPoint[(j-1+pointNum)%pointNum].y > polyPoint[j].y {
					p := new(XET)
					p.x = float64(polyPoint[j].x)
					p.ymax = polyPoint[(j-1+pointNum)%pointNum].y
					p.dx = float64(
						float64(polyPoint[(j-1+pointNum)%pointNum].x-polyPoint[j].x) /
							float64(polyPoint[(j-1+pointNum)%pointNum].y-polyPoint[j].y))
					p.next = pNET[i].next
					pNET[i].next = p
				}

				if polyPoint[(j+1+pointNum)%pointNum].y > polyPoint[j].y {
					p := new(XET)
					p.x = float64(polyPoint[j].x)
					p.ymax = polyPoint[(j+1+pointNum)%pointNum].y
					p.dx = float64(
						float64(polyPoint[(j+1+pointNum)%pointNum].x-polyPoint[j].x) /
							float64(polyPoint[(j+1+pointNum)%pointNum].y-polyPoint[j].y))
					p.next = pNET[i].next
					pNET[i].next = p
				}
			}
		}
	}

	for i := 0; i <= maxY; i++ {
		p := pAET.next
		for p != nil {
			p.x = p.x + p.dx
			p = p.next
		}
		tq := pAET
		p = pAET.next
		tq.next = nil
		for p != nil {
			for tq.next != nil && p.x >= tq.next.x {
				tq = tq.next
			}
			s := p.next
			p.next = tq.next
			tq.next = p
			p = s
			tq = pAET
		}

		q := pAET
		for q.next != nil {
			if q.next.ymax == i {
				q.next = q.next.next
			} else {
				q = q.next
			}
		}

		p = pNET[i].next
		q = pAET

		for p != nil {
			for q.next != nil && p.x >= q.next.x {
				q = q.next
			}
			s := p.next
			p.next = q.next
			q.next = p
			p = s
			q = pAET
		}

		p = pAET.next
		for p != nil && p.next != nil {
			for j := p.x; j <= p.next.x; j++ {
				paint(int(j), i)
			}
			p = p.next.next
		}
	}

}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 1200, 1000))
	pointList := []Point{{500, 100}, {200, 200}, {200, 700}, {500, 500}}
	PolyScan(pointList, 4, func(x, y int) {
		img.Set(x, y, color.RGBA{0, 0, 0, 255})
	})
	imgfile, _ := os.Create("test.png")
	defer imgfile.Close()
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}
