package main

import (
	"container/list"
	"crypto/sha1" // non-crypto hashes: adler32, crc32, crc64, fnv
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("strings:")
	fmt.Println(strings.Contains("test", "es"))               //true
	fmt.Println(strings.Replace("abcdabcdabcd", "a", "A", 2)) // AbcdAbcdabcd
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))        // a-b
	fmt.Println(strings.Split("a-b-c", "-"))                  // []string{"a", "b", "c"}

	fmt.Println("\nlists:")
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	fmt.Println("\nsort:")
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	fmt.Println("unsorted: ", kids)
	sort.Sort(ByName(kids))
	fmt.Println("sorted by name:", kids)
	sort.Sort(ByAge(kids))
	fmt.Println("sorted by Age:", kids)

	fmt.Println("\nhashing:")
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}
func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
