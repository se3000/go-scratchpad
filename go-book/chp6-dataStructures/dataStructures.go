package main

import "fmt"

func main() {
	fmt.Println("Arrays:")
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83
	var total float64 = 0

	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(total / float64(len(y)))

	total2 := 0.0
	z := [7]float64{13, 8, 5, 3, 2, 1, 1}
	for _, value := range z {
		total2 += value
	}
	fmt.Println(total2 / float64(len(z)))

	fmt.Println("Slices:")
	var j []float64
	var low = 5
	var high = 10
	j = make([]float64, low, high)

	//or
	arr := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	j = arr[low:high]
	fmt.Println("low:", low)
	fmt.Println("high:", high)
	fmt.Println("fullArray:", arr)
	fmt.Println("slice[low:high]", j)

	//both parameters are optional
	slice1 := arr[0:len(arr)]
	slice2 := arr[:]
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	slice3 := []int{1, 2, 3}
	slice4 := append(slice1, 4, 5)
	fmt.Println(slice3, slice4)

	slice5 := []int{1, 2, 3}
	slice6 := make([]int, 2)
	copy(slice6, slice5)
	fmt.Println("\n", slice5, slice6)

	fmt.Println("\n\nMaps:")
	map1 := make(map[string]int) //make initializes the object(?)
	map1["key1"] = 1
	map1["key2"] = 22
	fmt.Println(map1)
	delete(map1, "key2")
	fmt.Println(map1)
	fmt.Println(map1["key1"])
	fmt.Println(map1["key2"])

	if value, present := map1["key2"]; present {
		fmt.Println(value, present) //doesn't run
	}

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydorgen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "gas",
		},
	}
	fmt.Println("Easy syntax nested map:", elements)
}
