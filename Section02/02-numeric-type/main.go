package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64
var c int8 = 127 // -129, 128 will not run

func main() {
	programe05()
}

// programe01
func programe01() {
	a := 42
	b := 42.34325
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}

// programe02
func programe02() {
	x = 42
	y = 42.32345
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}

// This doesn't run
func programe03() {
	// x = 42.34234
	y = 42.32345
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}

// int8 example
func programe04() {
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}

// GOOS and GOARCH
func programe05() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
