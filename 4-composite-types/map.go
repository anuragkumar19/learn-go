package main

import (
	"fmt"
)

func main() {

	// A map is a reference to hash table
	// type --> map[K]V, where K is type of key, V is type of value, K must be comparable

	// Create new map
	a := make(map[string]string)

	a["Ram"] = "India"
	a["John"] = "America"

	b := map[string]int{
		"a": 4,
		"b": 5,
	}

	fmt.Println(a, b)
	fmt.Println(len(a), len(b))

	// Accessing map, we can access by using following syntax
	// If not exist it will return types zero value,
	// In case where we want to check for if it really existed or not we can use 2nd value usually called ok
	ramLocation, ok := a["Ram"]
	fmt.Println(ramLocation, ok)

	zeroLocation, ok := a["aaa"]
	fmt.Println(zeroLocation, ok)

	delete(a, "Ram")

	for name, loc := range a {
		fmt.Println(name+":", loc)
	}

	// We cannot take reference to any map value
	// zero value - nil , all func like len, delete and range loop are safe but storing to nil map cause panic

	// We can use map type as set type --> map[K]bool
	set := make(map[string]bool)

	set["A"] = true
	set["B"] = true
}
