package main

import "fmt" //compiler does not have to recomplie fmt each time
// import "go-book/chp11-packages/math"
import m "go-book/chp11-packages/math" // in case you want to use original math library

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
