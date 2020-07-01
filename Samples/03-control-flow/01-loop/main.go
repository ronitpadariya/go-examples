package main

import (
	"fmt"
)

func main() {
	programe02()
}

func programe01() {
	// for init; condition; post {}
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("The outer loop: %d\t The inner loop: %d\n", i, j)

		}
	}
}

func programe02() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}

func programe03() {
	x := 33
	for {
		if x > 122 {
			break
		}
		fmt.Printf("%v corresponds to %+q in ASCII\n", x, x)
		x++
	}
}
