package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	b := 1
	fmt.Println(b)
	b = 2
	fmt.Println(b)
	b = 100000
	fmt.Println(b)
	l := &b
	fmt.Println(l)
	*l = 100
	fmt.Println(b)
}