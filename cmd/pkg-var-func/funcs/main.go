package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return 
}

func main() {
	sum := add(1, 2)
	fmt.Println("Total = ", sum)
	//
	a,b := swap("Hoang", "Anh")
	fmt.Println("Swap = ", a, b)
	//
	fmt.Println(split(10))
}
