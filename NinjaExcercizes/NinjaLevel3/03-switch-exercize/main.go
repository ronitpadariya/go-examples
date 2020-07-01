package main

import (
	"fmt"
)

func main() {
	exe03()
}

func exe01() {
	switch {
	case true:
		fmt.Println("Should Print")
	case false:
		fmt.Println("Shouldn't Print")
	}
}

func exe02() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}

func exe03() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
