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

func mapConcurrentSplitedOrdered[T, U any](s []T, f func(T) U, splits ...int) []U {
	n_splits := 1
	if len(splits) > 0 {
		n_splits = splits[0]
	}

	elementsPerSplit := len(s) / n_splits
	leftover := len(s) % n_splits

	result := make([]U, len(s))

	var wg sync.WaitGroup

	for sp := 0; sp <= n_splits; sp++ {
		wg.Add(1)

		go func(sp int) {
			defer wg.Done()
			rest := 0
			if sp == n_splits {
				if leftover != 0 {
					rest = leftover
				} else {
					return
				}
			} else {
				rest = elementsPerSplit
			}
			for index := sp * elementsPerSplit; index < sp*elementsPerSplit+rest; index++ {
				result[index] = f(s[index])
			}

		}(sp)
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
func Map[T, U any](s []T, f func(T) U, splits ...int) []U {
	n_splits := 1
	if len(splits) > 0 {
		n_splits = splits[0]
	}
	if n_splits == 0 {
		return mapConcurrentOrdered(s, f)
	} else if n_splits == 1 {
		return mapSequentialOrdered(s, f)
	} else {
		return mapConcurrentSplitedOrdered(s, f, splits...)
	}
}

type mapOptions struct {
	Splits int
}

func NewMapOptions() *mapOptions {
	return &mapOptions{}
}

/*SetSequential
Squentially applies a function to each element in the slice.
*/
func (mo *mapOptions) SetSequential() {
	mo.Splits = 1
}

/*SetConcurrent
Applies the function to each element in the slice concurrently.
Each application will be processed on it's own goroutine.
*/
func (mo *mapOptions) SetConcurrent() {
	mo.Splits = 0
}

/*SetConcurrentSplits
Applies the function to each element in the slice concurrently.
The slice will be split in `n` different slices, and each slice will be
processed on it's own goroutine.
*/
func (mo *mapOptions) SetConcurrentSplits(n int) {
	mo.Splits = n
}
