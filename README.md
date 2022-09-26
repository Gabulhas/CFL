# Golang Concurrent and Functional Library.
A Golang implementation of Functional Language's functions like `Map`, `Filter`, `Fold`, but with a touch of concurrency in the mix.

## Use cases

Functional Programming like code is more readeable/clean and because some of the functions applied are [pure](https://en.wikipedia.org/wiki/Pure_function), we can apply these concurrently without locking.



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

    //Maps the square function to every element of the slice "myNumbers"
    result := cfl.Map(myNumbers, square, options)

    fmt.Println("My result:", result)
    //Outputs "My result: [1 4 9 16 25 36]"
}
```

## RoAdMaP
Implement every helper offered by [fp-go](https://github.com/repeale/fp-go), but making use of concurrency.

Also, make it so functions can return channels instead of the whole slice.

Checkout [Examples](/examples) to learn more about this project.
