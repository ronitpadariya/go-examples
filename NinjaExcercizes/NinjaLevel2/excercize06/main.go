package main

//	Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

import (
	"fmt"
)

const (
	a = 2020 + iota
	b = 2020 + iota
	c = 2020 + iota
	d = 2020 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
