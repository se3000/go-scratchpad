package main

import "fmt"

var outside string = "scope test"

// func outsideScope() {
// var other string = "NaNah Nah NaNa"
// }

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var y string
	y = "Hola World"
	fmt.Println(y)

	var z string
	z = "first"
	fmt.Println(z)
	z += " second"
	fmt.Println(z)
	fmt.Println(x == y)

	s := "Hello Go"
	fmt.Println(s)

	namingConvetion := "camelCase"
	fmt.Println("\nThe Golang naming convetion is", namingConvetion)

	fmt.Println(outside, "passed")
	// fmt.Println(other)

	const constant string = "\nI'm a const"
	fmt.Println(constant)
	// constant = "and another"

	var (
		a = "variables"
		b = "declared"
		c = "together"
	)
	fmt.Println(a, b, c)
}
