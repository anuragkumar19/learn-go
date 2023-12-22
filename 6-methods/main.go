package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

func main() {
	// OOP inspired, in Go an object is any value or variable which have methods
	// Methods - is a func associated with a particular type

	// Methods may be declared on any named type defined in the same package, so long as its underlying type is neither a pointer nor an interface.

	// In a realistic program, convention dictates that if any method of Point has a pointer receiver,
	// then all methods of Point should have a pointer receiver, even ones that don't strictly need it.

	p := Point{0, 0}
	q := Point{2, 2}

	// Methods are accessed same as any properties on type with dot(.) syntax
	fmt.Println(p.Distance(q))

	// Pointer receiver method
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	//If the receiver p is a variable of type Point but the method requires a *Point receiver, we can use this shorthand
	p.ScaleBy(2) // the compiler will perform an implicit &p on the variable

	// Point{2,2}.ScaleBy(4) // compile error no  way to obtain addr

	// We can call a Point method like Point.Distance with a *Point receiver, because there is
	// a way to obtain the value from the address: just load the value pointed to by the receiver. The
	// compiler inserts an implicit * operation for us.
	(&p).Distance(q)

	//  Nil Is a Valid Receiver Value - If a method accept nil we should tell in documentation

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	a := ColoredPoint{Point{1, 1}, red}
	b := ColoredPoint{Point{5, 4}, blue}

	// We can access methods of Point
	fmt.Println(p.Distance(a.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(b.Point)) // "10"

	/*
		A struct type may have more than one anonymous field
		then a value of this type would have all the methods of all types, and
		any addition al methods declared on type directly. When the compiler resolves a
		selector to a method, it first looks for a directly declared method name , then for methods promoted once from embedded fields, then for
		methods promoted twice from embedded fields, and so on. The compiler reports an error if the selector was ambiguous because two methods were promoted from
		the same rank

		Methods can be declared only on named types (like Point) and pointers to them (*Point),
		but thanks to embedding, it’s possible and sometimes useful for unnamed struct types to have
		methods too.
	*/

	var cache = struct {
		sync.Mutex
		mapping map[string]string
	}{
		mapping: make(map[string]string),
	}

	cache.Lock()
	cache.Unlock()

	/*
		The selector p.Distance yields a method value,a function that binds a method (Point.Distance) to a specific receiver value p
	*/
	distanceFromP := p.Distance        // method value
	fmt.Println(distanceFromP(q))      // "5"
	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)

	// Because of this methods can be passed as args to func like other func.

	/*
		Method expression
	*/
	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"
	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)            // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"

	/*
		Encapsulation

		A variable or method of an object is said to be encapsulated if it is inaccessible to clients of the
		object. Encapsulation, sometimes called information hiding, is a key aspect of object-oriented programming

		Encapsulation provides three benefits. First, because clients cannot directly modify the
		object’s variables, one need inspect fewer statements to understand the possible values of those variables

		Second, hiding implementation details prevents clients from depending on things that might
		change , which gives the designer greater freedom to evolve the implementation without breaking API compatibility.

		The third benefit of encapsulation, and in many cases the most important, is that it prevents
		clients from setting an object’s variables arbitrarily. Because the object’s variables can be set
		only by functions in the same package, the author of that package can ensure that all those
		functions maintain the object’s internal invariants.

		Functions that merely access or modify internal values of a type, such as the methods of the Logger type from log package,
		below, are called getters and setters. However, when naming a getter method, we usually omit the Get prefix
	*/

}

type Point struct {
	X, Y float64
}

type Path []Point

// Method declaration
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// Path is a named slice type, not a struct type like Point, yet we can still define methods for it.
// In allowing methods to be associated with any type
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// We can also receive pointer to the type
func (p *Point) ScaleBy(scale float64) {
	p.X *= scale
	p.Y *= scale
}

// Composing Types by Struct Embedding

// We can call methods of the embedded Point field using a receiver of type ColoredPoint, even though ColoredPoint has no declared methods
type ColoredPoint struct {
	Point
	Color color.RGBA
}
