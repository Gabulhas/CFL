package cfl

import (
	"sync"
)

// TODO: Improve this
// Appending is "slow", but seems unecessary to allocate such a big channel
func filterConcurrent[T any](s []T, f func(T) bool) []T {
	var wg sync.WaitGroup

	ch := make(chan T, len(s))

	wg.Add(len(s))
	go func() {
		for _, m := range s {
			go func(value T) {
				defer wg.Done()
				if f(value) {
					ch <- value
				}
			}(m)
		}

	}()
	wg.Wait()

	nFilteredValues := len(ch)

	result := make([]T, nFilteredValues)

	counter := 0
	for v := range ch {
		result[counter] = v
		counter = counter + 1
		if counter == nFilteredValues {
			return result
		}
	}
	//This should be an impossible case
	return []T{}

}

func filterConcurrentSplitted[T any](s []T, f func(T) bool, numberOfSplits int) []T {
	ranges := SplitRanges(len(s), numberOfSplits)

	var wg sync.WaitGroup
	ch := make(chan T, len(s))

	wg.Add(numberOfSplits)
	go func() {
		for _, r := range ranges {

			go func(r SplitRange) {
				defer wg.Done()
				for index := r.start; index < r.end; index++ {
					if f(s[index]) {
						ch <- s[index]
					}
				}
			}(r)
		}

	}()

	wg.Wait()
	nFilteredValues := len(ch)

	result := make([]T, nFilteredValues)

	counter := 0
	for v := range ch {
		result[counter] = v
		counter = counter + 1
		if counter == nFilteredValues {
			return result
		}
	}
	//This should be an impossible case
	return []T{}

}

func filterSequential[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Filter[T any](s []T, f func(T) bool, options ...*ConcurrencyOptions) []T {
	numberOfSplits := 1
	if len(options) > 0 {
		numberOfSplits = options[0].Splits
	} else {
		numberOfSplits = 1
	}
	if numberOfSplits == 0 {
		return filterConcurrent(s, f)
	} else if numberOfSplits == 1 {
		return filterSequential(s, f)
	} else {
		return filterConcurrentSplitted(s, f, numberOfSplits)
	}
}
