package main

import (
	"fmt"

	"github.com/Gabulhas/cfl"
)

// Returns the square of an integer value
func isEven(a int) bool {
	return a%2 == 0
}

//
func exampleIsEveryNumberEven() {
	myNumbers := []int{2, 4, 6, 8, 10, 12}

	result := cfl.Every(myNumbers, isEven, cfl.NewMapOptions().SetConcurrentSplits(3))
	fmt.Println("My result:", result)
	//My result: [1 4 9 16 25 36.....
}

func main() {
	exampleIsEveryNumberEven()
}
