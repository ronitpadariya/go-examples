package main

import (
	"fmt"
)

func main() {
	programe02()
}

// a new underlying array is allocated
func programe01() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x, 52, 53, 54, 55, 56, 57, 58, 59, 60) // new underlying array allocation

	fmt.Println(x)
	fmt.Println(y)
}

// the above array is used for TWO slices
func programe02() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x[:2], x[5:]...) // the same underlying array stores the value of the new slice

	fmt.Println(x)
	fmt.Println(y)
}
