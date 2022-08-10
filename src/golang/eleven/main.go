package main

import (
	"fmt"
	m "learner/src/golang/eleven/math"
	"math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
	fmt.Println(math.Pow(1.2, 4.5))
}
