package cfl

func Flat[T any](s [][]T) []T {

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
