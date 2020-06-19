package main

import (
	"fmt"
)

func main() {
	programe01()
}

//programe01
func programe01() {
	s := "Hello, world"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("---%x\n", "world")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("%#U \t %d", v, i)
	}
}
