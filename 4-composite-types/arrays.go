package main

import "fmt"

func main() {
	// Array- List of fixed size
	var arr [3]int

	fmt.Println(arr[0], arr[len(arr)-1])

	a := [4]int{1, 2, 3, 4}

	for _, x := range a {
		fmt.Println(x)
	}

	// `...` in arr literal - the number for element determined by initializer
	b := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("%T\n", b)

	// [3]int and [4]int are different types.

	// We can also specify index
	c := [...]string{0: "A", 1: "B", 2: "C", 3: "D"}

	fmt.Printf("%T, %s\n", c, c[0])

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "e", GBP: "g", RMB: "r"}
	fmt.Println(RMB, symbol[RMB])

	d := [...]int{99: 100}
	fmt.Println(len(d), d[0])

	// indices can be in any order some may be omitted

	e := [2]int{1, 2}
	f := [...]int{1, 2}
	g := [2]int{1, 3}
	// Comparable because for same type
	fmt.Println(e == f, g == e, f == g) // "true false false"

	// h := [3]int{1, 2}
	// fmt.Println(a == h) // compile error: cannot compare [2]int == [3]int
}
