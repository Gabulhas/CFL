# Golang Concurrent and Functional Library.
A Golang implementation of Functional Language's functions like `Map`, `Filter`, `Fold`, but with a touch of concurrency in the mix.

## Use cases

Functional Programming like code is more readeable/clean and because some of the functions applied are [pure](https://en.wikipedia.org/wiki/Pure_function), we can apply these concurrently.



## Installation

```bash
go get github.com/Gabulhas/cfl
```


## Example Usage

**Map** - Applies a function to every element in a slice.

```go
import (
        "fmt"

        "github.com/Gabulhas/cfl"
       )

// Returns the square of an integer value
func square(a int) int {
    return a * a
}

//
func exampleSquareNumbers() {
myNumbers := []int{1, 2, 3, 4, 5, 6}

           fmt.Println("My numbers:", myNumbers)

               options := cfl.NewMapOptions()
               options.SetConcurrentSplits(2)

               result := cfl.Map(myNumbers, square, options)
               fmt.Println("My result:", result)
}
```

Checkout [Examples](/examples) to learn more about this project.
