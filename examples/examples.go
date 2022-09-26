package main

import (
	"fmt"

	"github.com/Gabulhas/cfl"
)

var myNumbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

func isEven(a int) bool {
	return a%2 == 0
}

// Returns the square of an integer value
func square(a int) int {
	return a * a
}

func isLessThan15(a int) bool {
	return a < 15
}

func exampleEvery() {
	myNumbers := []int{2, 4, 6, 8, 10, 12}

	result := cfl.Every(myNumbers, isEven, cfl.NewMapOptions().SetConcurrentSplits(3))
	fmt.Println("[EVERY] Are my numbers less than 15:", result)
	//My result: [1 4 9 16 25 36.....
}

func exampleFilter() {
	result := cfl.Filter(myNumbers, isEven)
	fmt.Println("[FILTER] My even numbers:", result)
}

func exampleMap() {
	result := cfl.Map(myNumbers, square, cfl.NewMapOptions().SetConcurrentSplits(4))
	fmt.Println("[MAP] My numbers squared:", result)
}

func main() {
	exampleEvery()
	exampleFilter()
	exampleMap()
}
