package main

import (
	"fmt"
)

func main() {
	programe03()
}

func programe01() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}
}

// Map - add element & range
func programe02() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	m["Todd"] = 33

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}

}

// Map - delete
func programe03() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)

	fmt.Println(m["Miss MoneyPenny"])
	fmt.Println(m["Ian Fleming"])

	if v, ok := m["Miss MoneyPenny"]; ok {
		fmt.Println("value: ", v)
		delete(m, "MoneyPenny")
	}

	fmt.Println(m)

}
