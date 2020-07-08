package main

import (
	"fmt"
)

type vehicle struct {
	doors  int
	colors string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors:  2,
			colors: "white",
		},
		fourwheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors:  4,
			colors: "black",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)

}
