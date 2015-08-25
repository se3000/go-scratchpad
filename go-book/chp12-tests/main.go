package main

import "fmt"
import m "go-book/chp12-tests/math" // in case you want to use original math library

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
