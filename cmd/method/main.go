package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	//
	v1 := Vertex{3, 4}
	fmt.Println(Abs2(v1))
	//
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	//
	v3 := Vertex{3, 4}
	v3.Scale(10)
	fmt.Println(v3.Abs())
	//
	v4 := Vertex{3, 4}
	Scale(&v4, 10)
	fmt.Println(Abs(v4))
	//
	v5 := Vertex{3, 4}
	v5.Scale(2)
	ScaleFunc(&v5, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	///
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
