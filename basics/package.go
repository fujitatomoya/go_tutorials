package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Now you have %g problems.", math.Sqrt(7))

	// this leads to error, since pi is not exported
	// fmt.Println(math.pi)
	fmt.Println(math.Pi)
}
