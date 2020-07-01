package main

import (
	"fmt"
)

func main() {
	exe05()
}

func exe01() {
	for i := 1; i <= 10000; i++ {
		fmt.Printf("%v\n", i)
	}
}

func exe02() {
	for i := 65; i <= 95; i++ {
		fmt.Println(i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func exe03() {
	bd := 1986
	for bd <= 2020 {
		fmt.Println(bd)
		bd++
	}
}

func exe04() {
	bd := 1986
	for {
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}

func exe05() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the reminder (aka modulus) is %v\n", i, i%4)
	}
}
