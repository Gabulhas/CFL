package cfl

import (
	"fmt"
	"sync"
)

func flatConcurrent[T any](s [][]T) []T {
	var wg sync.WaitGroup
	var totalElements int

	for _, sp := range s {
		totalElements += len(sp)
	}
	result := make([]T, totalElements)

	wg.Add(len(s))

	go func() {
		lastStartIndex := 0
		for _, sp := range s {
			thisStartIndex := lastStartIndex
			go func(thisStartIndex int) {
				fmt.Println(thisStartIndex)
				localCounter := 0
				for _, value := range sp {
					result[localCounter+thisStartIndex] = value
					localCounter = localCounter + 1
				}
				wg.Done()
			}(thisStartIndex)
			lastStartIndex += len(sp)
		}

	}()

	wg.Wait()

	return result

}

func flatSequential[T any](s [][]T) []T {

	totalElements := 0

	for _, sp := range s {
		totalElements += len(sp)
	}

	result := make([]T, totalElements)

	counter := 0

	for _, sp := range s {
		for _, value := range sp {
			result[counter] = value
			counter = counter + 1
		}
	}
	return result

}

func Flat[T any](s [][]T, options ...*ConcurrencyOptions) []T {
	numberOfSplits := 1
	if len(options) > 0 {
		numberOfSplits = options[0].Splits
	} else {
		numberOfSplits = 1
	}

	switch numberOfSplits {
	case 1:
		return flatSequential(s)
	default:
		return flatConcurrent(s)
	}
}
