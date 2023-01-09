package main

import (
	"fmt"
	"math"
)

func main() {
	var n float32
	fmt.Scanf("%f", &n)
	ans := IEEE745(float64(n))
	fmt.Printf("%X\n", ans)
}

func IEEE745(n float64) int64 {
	if n == 0.0 {
		return 0
	}
	s := 0
	t := 0
	if n < 0.0 {
		s = 1
		n = math.Abs(n)
	}
	for {
		if n == math.Floor(n+0.5) {
			break
		}
		t--
		n = n * 2
	}
	x := int(math.Floor(n))
	for i := x; i > 1; i >>= 1 {
		t++
	}

	b := t
	if t < 0 {
		b = -t
	}
	m := x & ((1 << b) - 1)
	e := 127 + t
	ans := int64((s << 31) + (e << 23) + (m << (23 - b)))
	return ans
}
