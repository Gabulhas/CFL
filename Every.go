package cfl

import "sync"

func everyConcurrent[T any](s []T, f func(T) bool) bool {

	var wg sync.WaitGroup
	ch := make(chan bool, 10)

	go func() {
		wg.Add(len(s))
		for _, m := range s {
			go func(value T) {
				defer wg.Done()
				if !f(value) {
					ch <- false
				}
			}(m)
		}

		wg.Wait()
		close(ch)
	}()

	for range ch {
		return false
	}
	return true

}
func everyConcurrentSplitted[T any](s []T, f func(T) bool, numberOfSplits int) bool {
	ranges := SplitRanges(len(s), numberOfSplits)

	var wg sync.WaitGroup
	ch := make(chan bool, 10)

	go func() {
		for _, r := range ranges {
			wg.Add(1)

			go func(r SplitRange) {
				defer wg.Done()
				for index := r.start; index < r.end; index++ {
					if !f(s[index]) {
						ch <- false
					}
				}
			}(r)
		}
		wg.Wait()
		close(ch)

	}()

	for range ch {
		return false
	}
	return true

}

func everySequencial[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func Every[T any](s []T, f func(T) bool, options ...*ConcurrencyOptions) bool {
	numberOfSplits := 1
	if len(options) > 0 {
		numberOfSplits = options[0].Splits
	} else {
		numberOfSplits = 1
	}
	if numberOfSplits == 0 {
		return everyConcurrent(s, f)
	} else if numberOfSplits == 1 {
		return everySequencial(s, f)
	} else {
		return everyConcurrentSplitted(s, f, numberOfSplits)
	}

}
