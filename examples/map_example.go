package main

import (
	"fmt"

	"github.com/Gabulhas/cfl"
)

// Returns the square of an integer value
func square(a int) int {
	return a * a
}

//
func exampleSquareNumbers() {
	myNumbers := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("My numbers:", myNumbers)

	options := cfl.NewMapOptions()
	options.SetConcurrentSplits(2)

	result := cfl.Map(myNumbers, square, options)
	fmt.Println("My result:", result)
}

func main() {
	exampleSquareNumbers()
}
