package main

import "fmt"

func SliceExamples() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, 2, 3, 1}

	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
}
