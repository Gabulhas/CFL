package cfl

import "sync"

func mapConcurrentOrdered[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))

	var wg sync.WaitGroup

	wg.Add(len(s))

	for i, m := range s {
		go func(index int, value T) {
			defer wg.Done()
			result[index] = f(value)
		}(i, m)
	}
	wg.Wait()
	return result
}

func mapConcurrentSplitedOrdered[T, U any](s []T, f func(T) U, numberOfSplits int) []U {
	ranges := SplitRanges(len(s), numberOfSplits)

	result := make([]U, len(s))

	var wg sync.WaitGroup

	for _, r := range ranges {
		wg.Add(1)

		go func(r SplitRange) {
			defer wg.Done()
			for index := r.start; index < r.end; index++ {
				result[index] = f(s[index])
			}
		}(r)
	}
	wg.Wait()
	return result
}

func mapSequentialOrdered[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for index, value := range s {
		result[index] = f(value)
	}
	return result
}

/*Map
Orderly Maps a function over every element in a slice
s is the slice
f is the function
options is an MapOptions
*/
func Map[T, U any](s []T, f func(T) U, options ...*ConcurrencyOptions) []U {
	numberOfSplits := 1
	if len(options) > 0 {
		numberOfSplits = options[0].Splits
	} else {
		numberOfSplits = 1
	}
	if numberOfSplits == 0 {
		return mapConcurrentOrdered(s, f)
	} else if numberOfSplits == 1 {
		return mapSequentialOrdered(s, f)
	} else {
		return mapConcurrentSplitedOrdered(s, f, numberOfSplits)
	}
}
