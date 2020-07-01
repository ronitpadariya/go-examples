package main

import (
	"fmt"
)

func main() {
	exe02()
}

func exe01() {
	x := "Bond"

	if x == "Bond" {
		fmt.Printf("Bond, James Bond")
	}
}

func exe02() {
	x := "MoneyMoney"

	if x == "MoneyMoney" {
		fmt.Println(x)
	} else if x == "Jaames Bond" {
		fmt.Println("BONDONDONDOND", x)
	} else {
		fmt.Println("neither nor")
	}
}
