package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	var v *Vertex = new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}
