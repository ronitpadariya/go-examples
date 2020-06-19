package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)

func main() {
	programe03()
}

func programe01() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

func programe02() {
	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)
}

const (
	d0 = iota
	d1
	d2
)

func programe03() {
	fmt.Println(d0)
	fmt.Println(d1)
	fmt.Println(d2)
}
