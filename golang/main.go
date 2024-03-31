package main

import (
	"fmt"
	"math"
	"math/rand"
)

func doit(total int, res chan int) {
	inside := 0
	var x float64
	var y float64

	for i := 0; i < total; i++ {
		x = rand.Float64()
		y = rand.Float64()
		if x*x+y*y < 1 {
			inside += 1
		}
	}
	res <- inside
}

func main() {
	inside := 0
	total := 10000000000
	buckets := 100
	h := total / buckets

	summen := make(chan int)
	for i := 0; i < h; i++ {
		go doit(h, summen)
	}
	fmt.Println("Started all rutines. Now waiting for results")
	for i := 0; i < h; i++ {
		inside += <-summen
	}
	var my_pi float64
	my_pi = 4.0 * float64(inside) / float64(total)
	err_co := math.Abs(math.Pi - float64(my_pi))
	fmt.Println(my_pi, err_co)
}
