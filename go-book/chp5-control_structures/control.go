package main

import "fmt"

func main() {
	fmt.Println("for:")
	var i int
	for i < 10 {
		i += 1
		fmt.Println(i)
	}

	fmt.Println("\nif:")
	j := 1
	for j <= 4 {
		if j%2 == 0 {
			fmt.Println(j, "even")
		} else {
			fmt.Println(j, "odd")
		}
		j++ //++ works!
	}

	fmt.Println("\nswitch:")
	l := 0
	for l < 5 {
		switch l {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		}
		l++
	}
}
