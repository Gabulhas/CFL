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
	myNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	result := cfl.Map(myNumbers, square, cfl.NewMapOptions().SetConcurrentSplits(4))
	fmt.Println("My result:", result)
	//My result: [1 4 9 16 25 36.....
}
