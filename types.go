package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1+1.0)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println("Hello ", "World")
	fmt.Println("\nBooleans:\n")
	fmt.Println(true && false)
	fmt.Println(false || true)
	fmt.Println(!true)
}
