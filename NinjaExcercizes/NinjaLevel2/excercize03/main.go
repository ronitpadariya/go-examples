package main

// Create TYPED and UNTYPED constants. Print the values of the constants.

import (
	"fmt"
)

const (
	a     = 42
	b int = 43
)

func main() {
	fmt.Println(a, b)
}
