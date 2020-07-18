package main

import (
	"fmt"
)

// Method set
func main() {
	x := 0
	fmt.Println("X before ", &x)
	fmt.Println("X before  ", x)
	foo(&x)
	fmt.Println("X after ", &x)
	fmt.Println("X after ", x)
}

func foo(y *int) {
	fmt.Println("y before ", y)
	fmt.Println("y before ", *y)
	*y = 43
	fmt.Println("y after ", y)
	fmt.Println("y after ", *y)
}
