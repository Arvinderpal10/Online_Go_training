package main

import (
	"fmt"
)

// var t int = 10
// var s string = 100.50
// var u bool = true

// group decalaration
// var (
//
//	t int = 10
//	s string = 100.50
//	u bool = true
//	)

// Shadowing

// global

var x int = 200

func main() {
	fmt.Println("Varaibles")
	// x := math.Min(23, 409)
	x = 300
	fmt.Println(x)
	//   var v int = 100
	// var c = 100
	//  c := 100
	// Multiple variable decalaration
	// a, b := 10, 20
	var a, b, c = 10, 200.56, "Hello"
	fmt.Println(a, b, c)

}
