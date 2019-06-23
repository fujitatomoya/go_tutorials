package main

import (
	"fmt"
	"math"
)

// receiver structure for methods,
type MyFloat float64
type Vertex struct {
	X, Y float64
}

//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	fmt.Println(f)
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
