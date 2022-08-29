package main

import (
	"fmt"

	"github.com/Gabulhas/cfl"
)

//
func exampleSquareNumbers() {
	myNumbers := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("My numbers:", myNumbers)

	squareNumbersFunc := func(a int) int {
		return a * a
	}
	options := cfl.NewMapOptions()
	options.SetConcurrentSplits(2)

	result := cfl.Map(myNumbers, squareNumbersFunc, options)
	fmt.Println("My result:", result)
}

func main() {
	exampleSquareNumbers()
}
