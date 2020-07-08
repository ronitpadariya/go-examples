package main

import (
	"fmt"
)

type person struct {
	first       string
	last        string
	favFlavours []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavours: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavours: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavours {
			fmt.Println(i, val)
		}
		fmt.Println("----------------")
	}

}
