package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	aname := names[0:2]
	bname := names[1:3]
	fmt.Println(aname, bname)

	bname[0] = "XXX" // slices are like references to array.
	fmt.Println(aname, bname)
	fmt.Println(names)
}
