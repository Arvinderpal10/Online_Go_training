package main

import "fmt"

func IncreasingNumber(n int) int {
	if n >= 5 {
		return 1
	}
	fmt.Println(n)
	return IncreasingNumber(n + 1)

}

// func first() {
// 	fmt.Println("I am in first function")
// 	third()
// 	fmt.Println("Going to main")

// }
// func second() {
// 	fmt.Println("I am in Second function")

// }
// func third() {

// 	fmt.Println("I am in Third function")
// 	second()
// 	fmt.Println("Inside third")

// }

// func add(a int, b int) int {
// 	result := a + b
// 	// fmt.Println(result)
// 	return result
// 	// fmt.Println("This will not run")
// }

//  Function Body
// func Greet() {
// 	fmt.Println("Hello World")
// }

// user defined function
// func --
// main() , Println(), Scan()
func main() {
	// fmt.Println("Hello World")
	// fmt.Println("Hello World")
	// fmt.Println("Hello World")
	// fmt.Println("Hello World")
	// fmt.Println("Hello World")
	// Greet()
	// Greet()  // Calling a function
	// for loops

	//  create a function and call it five times

	// parameters--> We are using them in function definiton  ---
	// Aruguments -- Passed while calling the function
	// add(24, 20)
	// add(204, 120)
	// result := add(20, 120)
	// result += 10
	// fmt.Println(result)
	// return --

	// Without arguments without return
	// func Greet() {
	// 	fmt.Println("Hello World")
	// }

	//  with arguments without return
	// func add(a int, b int){
	// 	result := a + b
	// 	// fmt.Println(result)
	// }

	// without argumnets with return
	// func add() int {
	//  var a int = 10
	// var b int = 40
	// 	result := a + b
	// 	return result

	// }

	// with argumnets with return

	// func add(a int, b int) int {
	// 	result := a + b
	// 	// fmt.Println(result)
	// 	return result
	// }

	//
	// fmt.Println("Starting")
	// first()
	// fmt.Println("Ended")

	//  recursion -- Self calling  function

	fmt.Println(IncreasingNumber(0))
	// fmt.Println(IncreasingNumber(4))

}
