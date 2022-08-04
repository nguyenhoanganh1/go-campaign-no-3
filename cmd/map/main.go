package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func make_map() {
	m = make(map[string]Vertex)
	m["Bel Labs"] = Vertex{
		30.1123,
		1233.412,
	}
	fmt.Println(m)
}

func map_liter() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}
func map_liter2() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func mutating_maps() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	make_map()
	map_liter()
	map_liter2()
	mutating_maps()

}
