package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// func main() {
// 	// This is comment

// 	var s string

// 	argLen := len(os.Args)

// 	for i := 1; i < argLen; i += 1 {
// 		s += os.Args[i]

// 		if i != argLen-1 {
// 			s += " "
// 		}
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	var s, sep string

// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}

// 	fmt.Println(s)
// }

// func main() {
// 	// Infinite loop,break/return can be used to break
// 	// for {}

// 	// While loop
// 	// for condition {}

// 	s, sep := "", ""
// 	// Range
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = " "
// 	}

// 	fmt.Println(s)
// }

// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// }

// func main() {
// 	fmt.Println(os.Args[1:])
// }

func main() {
	// Exercise 1.1
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("Join", time.Since(start).Microseconds())

	start2 := time.Now()
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
	fmt.Println("Loop", time.Since(start2).Microseconds())

	// Exercise 1.2
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}

	// Exercise 1.3

}
