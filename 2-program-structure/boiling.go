// Boiling prints the boiling point of water
package main

import "fmt"

/*
	Package level declaration -

The name of each package-level entity is visible
not only throughout the source file that contains its declaration,
but throughout all the files of the package
*/
const boilingPoint = 212.0

func main() {
	var f = boilingPoint
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
