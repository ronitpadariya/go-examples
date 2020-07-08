package main

import (
	"fmt"
)

// we create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS

var x int

type person struct {
	first string
	last  string
}

type foo int

var y foo

const bar = 42

func main() {
	y = bar
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)
}
