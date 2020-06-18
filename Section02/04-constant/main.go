package main

import (
	"fmt"
)

const (
	a int     = 42
	b float64 = 42.78
	c string  = "James Bond"
)

const x = 42
const y float64 = 42.2

type hotdog int
type hotcat float64

func main() {
	programe02()
}

// programe01
func programe01() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

// programe02
func programe02() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Printf("%T\n", hotdog(x))
	fmt.Printf("%T\n", hotdog(y))

	fmt.Printf("%T\n", hotcat(x))
	fmt.Printf("%T\n", hotcat(y))
}
