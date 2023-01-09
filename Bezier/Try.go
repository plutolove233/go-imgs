package main

import "fmt"

func try(t *[]int) {
	*t = append(*t, 10)
}

func main() {
	var a []int
	try(&a)
	fmt.Println(a)
}
