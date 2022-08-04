package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var c, python, java bool

var x, y int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	var i int
	fmt.Println(i, c, python, java)
	//
	var d, py, ja = true, false, "no!"
	fmt.Println(x, y, d, py, ja)
	//
	var u, o int = 3, 4
	k := 3
	c, p, j := true, false, "no!"
	fmt.Println(u, o, k, c, p, j)
	//
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	//
	var e int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", e, f, b, s)
	//
	var kk, l int = 3, 4
	var pop float64 = math.Sqrt(float64(kk*kk + kk*l))
	var uu uint = uint(pop)
	fmt.Println(kk, l, uu)
	// type interface
	vv := 32
	fmt.Println("v is of type %T\n", vv)
	// const
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
	//
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
