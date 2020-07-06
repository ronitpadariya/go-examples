package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazwlnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	for i, jm := range xp {
		fmt.Println("record: ", i)
		for j, val := range jm {
			fmt.Printf("\t index position %v \t value: \t %v \n", j, val)
		}
	}
}
