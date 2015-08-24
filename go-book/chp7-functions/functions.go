package main

import "fmt"

func main() {
	xxs := []float64{98, 97, 77, 82, 83}
	fmt.Println(average(xxs))
	fmt.Println(named())
	fmt.Println(nonNamed())
	x1, x2 := multipleValues()
	fmt.Println(x1, x2)
	fmt.Println("vadiadic add(1, 2, 3):", add(1, 2, 3))
	xs := []int{1, 2, 3}
	fmt.Println("convert slice to many args:", add(xs...))

	demonstrateClosure()

	nextEven := makeEvenGenerator()
	fmt.Println("next even:", nextEven())
	fmt.Println("next even:", nextEven())
	fmt.Println("next even:", nextEven())

	fmt.Println("\nfactorial", factorial(5))

	demonstateDefer()

	panicAndRecover()
}

func average(xs []float64) float64 {
	// panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func named() (returnValue int) {
	returnValue = 1
	fmt.Println("named function:")
	return
}

func nonNamed() (returnValue int) {
	returnValue = 1
	fmt.Println("named function psyche out:")
	return 3
}

func multipleValues() (string, string) {
	fmt.Println("multivalue return:")
	return "(first value: often answer),", "(second value: often error/success||failure)"
}

//Variadic Functions
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func demonstrateClosure() {
	x := 0

	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func demonstateDefer() {
	defer fmt.Println("defered til end of function: Second\n")
	fmt.Println("\nnot defered: first")

	// use case opening a file:
	//
	// f, _ := os.Open(filename)
	// defer f.Close()
	//
	// ensures the file closes
}

func panicAndRecover() {
	defer func() {
		str := recover()
		fmt.Println("RECOVERED:", str)
	}() //anonymous function
	panic("PANIC")
}
