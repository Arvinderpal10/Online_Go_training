package main

import "fmt"

func main() {
	fmt.Println("Whitespace")
	var x int = 100
	fmt.Println("      New       variable:", x)

	// Escape Sequences
	fmt.Println("Hello \n", x) // NEW LINE
	fmt.Println("Hello\t", x)
	fmt.Println("Hello \" Again \" ")

	// const MAX = 100
	// MAX = 200
	// fmt.Println(MAX)
	// var y int = 23
	// y = 40
	// fmt.Println(y)

	// Operators -- Arithmetic  -- + , - , / , % , *
	fmt.Println(20 + 40)
	fmt.Println(30 + 20.00) // Explicit Conversion
	fmt.Println(30 / 5)
	fmt.Println(float32(30) / 4) //
	// % - Modulus Operator - Gives the remainder  5 %2  == 1
	fmt.Println(30 % 4)
	fmt.Println(30 * 5)
	// fmt.Println(30.00 * 5)

	var u int = 20
	u += 30 // u = u + 30

	// u -=10   // u = u -10
	// u *= 20
	// u /=5
	fmt.Println(u)

}
