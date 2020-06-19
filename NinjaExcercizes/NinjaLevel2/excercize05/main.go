package main

// Create a variable of type string using a raw string literal. Print it.

import (
	"fmt"
)

func main() {
	a := `here is something
	as 
	a 
	raw string
	literal
	"you see"
	another thing`
	fmt.Println(a)
}
