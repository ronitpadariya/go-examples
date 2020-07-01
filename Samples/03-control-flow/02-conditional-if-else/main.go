package main

import (
	"fmt"
)

func main() {
	programe04()
}

func programe01() {
	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 2 == 2 {
		fmt.Println("005")
	}

}

func programe02() {

	if x := 42; x == 42 {
		fmt.Println("001")
	}
}

func programe03() {
	x := 42
	if x == 40 {
		fmt.Println("our value was 40")
	} else {
		fmt.Println("our value was not 40")
	}
}

func programe04() {
	x := 434
	if x == 40 {
		fmt.Println("our value was 40")
	} else if x == 41 {
		fmt.Println("our value was 41")
	} else {
		fmt.Println("our value was not 40 or 41")
	}
}
