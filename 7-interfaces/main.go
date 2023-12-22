package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		Interface in Go are abstract type. They don't define exact shape of type. They only define gew methods of type
		In Go, interfaces are satisfied implicitly, we don't have tell which-which interface a type satisfies

		An interface type specifies a set of methods that a concrete type must possess to be considered
		an instance of that interface.
	*/
	// io.Writer a interface with this definition
	/*
		The io.Writer type is one of the most widely used interfaces because it provides an abstraction of all the types to which bytes can be written, which includes files, memory buffers, network connections, HTTP clients, archivers, hashers, and so on. The io package defines many other useful interfaces.

		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	*/

	// This mean any type with a method Write with the same signature will satisfy the interface

	var c ByteCounter

	// func fmt.Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	// Because ByteCounter satisfies io.Writer we can pass it in fmt.Fprintf
	fmt.Fprintf(&c, "Hello Writer")

	fmt.Println(c)

	/*
		Interface satisfaction:

		A type satisfies an interface if it possesses all the methods the interface requires

		The assignability rule for interfaces is very simple: an expression may be assigned to an interface only if its type satisfies the interface.
	*/
	var w io.Writer
	w = os.Stdout         // OK: *os.File has Write method
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	// w = time.Second       // compile error: time.Duration lacks Write method
	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	// rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
	fmt.Fprintf(w, "")
	fmt.Fprintf(rwc, "")

	w = rwc // OK: io.ReadWriteCloser has Write method
	// rwc = w // compile error: io.Writer lacks Close method

	/*
		NOTE: each named concrete type T, some of its methods have a receiver of type T itself where as others require a *T pointer. Recall also that it is legal to
		call a *T method on an argument of type T so long as the argument is a variable; the compiler implicitly takes its address. But this is mere syntactic sugar : a value of typ e T does not possess
		all the methods that a *T pointer does, and as a result it might satisfy fewer interfaces.
	*/
	// var _ = IntSet{}.String() // compile error: String requires *IntSet receiver

	var s IntSet
	var _ = s.String() // OK: s is a variable and &s has a String method

	var _ fmt.Stringer = &s // OK
	// var _ fmt.Stringer = s  // compile error: IntSet lacks String method

	/*
		An interface wraps and conceals the concrete type and value that it holds. Only the methods revealed by the interface type may
		be called, even if the concrete type has others
	*/
	os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
	os.Stdout.Close()                // OK: *os.File has Close method
	var w2 io.Writer
	w2 = os.Stdout
	w2.Write([]byte("hello")) // OK: io.Writer has Write method
	// w2.Close()                // compile error: io.Writer lacks Close method

	// interface {}   === any
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type IntSet struct { /* ... */
}

func (*IntSet) String() string
