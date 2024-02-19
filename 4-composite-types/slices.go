package main

import "fmt"

func main() {
	// Slices are variable length list. They are pointer to an underlying array.
	// Slice contain three info:- ptr, len, cap
	// ptr - to first element of slice. len - length of slice, cap- capacity

	a := [...]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}

	// Weekend is a slice which points to underlying array `a`
	// ptr - points to "SAT"(i=5), len=2,cap=2
	weekend := a[5:7]
	fmt.Printf("type:%T, len: %d, cap: %d\n", weekend, len(weekend), cap(weekend))
	fmt.Println(weekend)

	// Workday is a slice which points to underlying array `a`
	// ptr - points to "MON"(i=0), len=5,cap=7
	workday := a[:5]
	fmt.Printf("type:%T, len: %d, cap: %d\n", workday, len(workday), cap(workday))
	fmt.Println(workday)

	// As they are pointer passing then in func, func can modify underlying array
	reverse(a[:])
	fmt.Println(a)

	// Reset
	reverse(a[:])

	// directly create a slice
	b := []string{"JAN", "FEB"}
	fmt.Printf("type:%T, len: %d, cap: %d\n", b, len(b), cap(b))

	// Slice are not comparable
	// Only valid comparison is ---- slice != nil or slice == nil
	// len(nil) = 0 & len([]int{}) = 0 so, as a standard we should always check ---- len(slice) != 0

	// Append
	// Always reassign after append
	b = append(b, "MAR")
	fmt.Printf("type:%T, len: %d, cap: %d\n", b, len(b), cap(b))

	b = append(b, "APR")
	fmt.Printf("type:%T, len: %d, cap: %d\n", b, len(b), cap(b))

	b = append(b, "MAY")
	fmt.Printf("type:%T, len: %d, cap: %d\n", b, len(b), cap(b))

	// By above results we can see how cap vary

	// We can now implement our own appendString
	appendString(b, "JUN")
	fmt.Printf("type:%T, len: %d, cap: %d\n", b, len(b), cap(b))

	// We can use make func to create slice of given (type, len ,cap)
	c := make([]int, 5, 10)
	fmt.Printf("type:%T, len: %d, cap: %d\n", c, len(c), cap(c))
}

// Reverse a slice of int
func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendString(x []string, y string) []string {
	var z []string
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]string, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
