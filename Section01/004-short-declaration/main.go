package main

import "fmt"

var a int
var b string = "Steve Roger"
var c bool
var d bool = true

func main() {
	e := 42
	f := "Shaken not stirred"
	g := `Miss Moneypenny says, "Helloooooo Steve."`
	h := `Q says, "I have some toys for you, Steve."`

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(b, "says,", f)
	fmt.Println(g)
	fmt.Println(h)
}
