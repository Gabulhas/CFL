package cfl

type ConcurrencyOptions struct {
	Splits int
}

func NewMapOptions() *ConcurrencyOptions {
	return &ConcurrencyOptions{}
}

/*SetSequential
Squentially applies a function to each element in the slice.
*/
func (mo *ConcurrencyOptions) SetSequential() *ConcurrencyOptions {
	mo.Splits = 1
	return mo
}

/*SetConcurrent
Applies the function to each element in the slice concurrently.
Each application will be processed on it's own goroutine.
*/
func (mo *ConcurrencyOptions) SetConcurrent() *ConcurrencyOptions {
	mo.Splits = 0
	return mo
}

/*SetConcurrentSplits
Applies the function to each element in the slice concurrently.
The slice will be split in `n` different slices, and each slice will be
processed on it's own goroutine.
*/
func (mo *ConcurrencyOptions) SetConcurrentSplits(n int) *ConcurrencyOptions {
	mo.Splits = n
	return mo
}
