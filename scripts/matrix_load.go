package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// Allocate a zeroed real matrix of size 3×5
	zero := mat.NewDense(3, 5, nil)

	fmt.Println(zero)
}
