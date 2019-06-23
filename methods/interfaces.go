package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var j interface{}
	describe2(j)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}