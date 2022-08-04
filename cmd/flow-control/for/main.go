package main

import "fmt"

func sum() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sum_con() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

func while() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
func forever(){
	for {

	}
}
func main() {
	sum()
	sum_con()
	while()
	forever()
}
