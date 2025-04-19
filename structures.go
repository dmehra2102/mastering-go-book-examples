package main

import "fmt"

type xyz struct {
	x int
	y int
	z int
}

type zyx struct {
	z int
	y int
	x int
}

func GoStructure() {
	p1 := xyz{2, 1, 4}
	p2 := zyx{x: 2, z: 4, y: 1}

	fmt.Print(p1)
	fmt.Print(p2)
}
