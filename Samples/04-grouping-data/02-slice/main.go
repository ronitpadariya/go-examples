package main

import (
	"fmt"
)

func main() {
	// a SLICE allows you to group together VALUES of the same TYPE
	// Slice - composite literal
	programe01()

	// Slice - for range
	// programe02()

	// Slice - slicing a slice
	// programe03()

	// Slice - append to a slice
	// programe04()

	// Slice - deleting from a slice
	// programe05()

	// Slice - make
	// programe06()

	// Slice - multi-dimensional slice
	// programe07()
}

// a SLICE allows you to group together VALUES of the same TYPE
// Slice - composite literal
func programe01() {
	// COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
}

// Slice - for range
func programe02() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Println(i, v)
	}
}

// Slice - slicing a slice
func programe03() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}

// Slice - append to a slice
func programe04() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)
}

// Slice - deleting from a slice
func programe05() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

// Slice - make
func programe06() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[1] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3425)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

// Slice - multi-dimensional slice
func programe07() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazwlnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
