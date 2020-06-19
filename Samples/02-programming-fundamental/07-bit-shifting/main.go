package main

import (
	"fmt"
)

func main() {
	programe05()
}

// programe01
func programe01() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
}

// programe02
func programe02() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

// programe03
const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func programe03() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

// programe04
func programe04() {
	s := 6
	fmt.Printf("%d\t\t%b\n", s, s)
	s = s << 1
	fmt.Printf("%d\t\t%b\n", s, s)
	s = s << 1
	fmt.Printf("%d\t\t%b\n", s, s)

	kb := 1024
	mb := 1024 * 1024
	gb := 1024 * 1024 * 1024
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

// programe05
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func programe05() {
	fmt.Println("binary\t\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t", GB)
	fmt.Printf("%d\n", GB)
}
