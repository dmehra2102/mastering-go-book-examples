package function

import (
	"fmt"
	"os"
	"strconv"
)

// Anonymous functions
var anonyFunc = func() {
	fmt.Println("I'm a anonymous function.")
}

// Functions that return multiple values
func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

// The return values of a function can be named!
func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

func FunctionMain() {
	anonyFunc()
	min, max := namedMinMax(3, -4)
	fmt.Println(min, max)

	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("The program needs 1 argument!")
		return
	}
	y, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	square := func(s int) int {
		return s * s
	}
	fmt.Println("The square of", y, "is", square(y))
	double := func(s int) int {
		return s + s
	}
	fmt.Println("The double of", y, "is", double(y))
	fmt.Println(doubleSquare(y))
	d, s := doubleSquare(y)
	fmt.Println(d, s)

}
