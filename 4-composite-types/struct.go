package main

import (
	"fmt"
)

// Struct are aggregate data type grouping together zero or more fields and value of arbitrary data type
type User struct {
	ID           int
	Name         string
	Username     string
	Email        string
	FriendsCount int
}

func main() {

	// Fields have zero value by default
	var user User

	// Dot notation to access value
	user.ID = 1
	user.Name = "Anurag"
	user.Username = "anurag"
	user.Email = "ANURAG@EXAMPLE.COM"
	user.FriendsCount = 0

	fmt.Println(user)

	userPtr := &user
	// We can also user dot notation with ptr
	fmt.Println(userPtr.Name)

	willNotModifyUser(user)
	fmt.Println(user)

	willModifyUser(&user)
	fmt.Println(user)

	// Struct name & fields follow exporting case convention
	type Point struct {
		X, Y int // Shorthand syntax
		z    int // Not accessible outside pkg
	}

	// A struct cannot contain same struct type as field type. But it can have ptr to its own type
	type Tree struct {
		val   int
		left  *Tree
		right *Tree
	}

	// Struct literal
	p := Point{2, 3, 5}
	user2 := User{
		ID:   2,
		Name: "Hello",
	}

	fmt.Println(p, user2)

	// We can also directly get ptr
	user3 := &User{}
	fmt.Println(user3.ID)

	// If all fields of struct are comparable struct is comparable, comparable struct can be used as key in map
	fmt.Println(user == user2)

	// Struct embedding & anonymous fields

	// 1.
	// type Circle struct {
	// 	Center Point
	// 	Radius int
	// }

	// type Wheel struct {
	// 	Circle Circle
	// 	Spokes int
	// }
	// for 1.
	// var wheel Wheel
	// fmt.Println(wheel.Circle.Center.X, wheel.Circle.Center.Y)

	// or 2.
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}

	var wheel Wheel

	wheel.X = 1
	wheel.Y = 2
	wheel.Radius = 10

	fmt.Println(wheel.X, wheel.Y) // we can access fields directly using struct embedding
	// wheel.Point.X & .Y are still valid
	fmt.Println(wheel.Point.X, wheel.Point.Y)

	// but we cannot use this with struct literal
	wheel = Wheel{
		Circle: Circle{
			Point: Point{
				X: 1,
				Y: 2,
			},
		},
		// X: 2, // compile err
	}

	fmt.Printf("%v\n", user)

	// We can user # adverb with printf %v to log struct as struct literal
	fmt.Printf("%#v\n", user)
}

func willNotModifyUser(user User) {
	user.Name = "Ram"
}

func willModifyUser(user *User) {
	user.Name = "Ram"
}
