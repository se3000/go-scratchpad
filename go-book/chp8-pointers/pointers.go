package main

import "fmt"

func main() {
	x := 5
	setZeroPassByValue(x)
	fmt.Println("zero? :", x)

	// &someVar references a pointer to someVar
	setZeroPassByReference(&x)
	fmt.Println("zero? :", x)

	//other way to get a pointer
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}

func setZeroPassByValue(x int) {
	x = 0
}

func setZeroPassByReference(xPtr *int) {
	// *someVar dereferences a pointer to someVar
	*xPtr = 0
}

func one(x *int) {
	*x = 1
}
