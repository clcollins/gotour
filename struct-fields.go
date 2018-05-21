package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.Y = 5
	fmt.Printf("Y: %v\n", v.Y)
}
