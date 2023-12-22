package main

import "fmt"

func main() {
	// var name type = expression
	var s string = "string"
	var integer = 12

	// Tuple assignment
	var i, j = 1, 2
	var a, b int
	var p, q, r = true, "hello", 4

	// Shorthand decl

	// One subtle but important point:a short variable declaration does not necessarily declare all the
	// variables on its left-hand side. If som e of them were already declared in the same lexical block
	// then the short variable declaration acts like an assignment to those variables

	// A short variable declaration must declare at least one new variable, however, so this code will
	// not compile

	fmt.Println(s, integer, i, j, a, b, p, q, r)
}
