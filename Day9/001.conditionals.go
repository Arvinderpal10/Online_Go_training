package main

import "fmt"

func main() {
	// fmt.Println("Second Line")
	// fmt.Println("Hello")

	// Conditionals helps us to make decision in the code

	// if condition {
	// 	// Run this if condition is true
	// }

	// if true {
	// 	fmt.Println("Hello")
	// }
	// if false {
	// 	fmt.Println("Hello")
	// } else {
	// 	fmt.Println("Inside Else")
	// }

	// age := 19
	// if age > 18 {
	// 	fmt.Println("Elligible to vote")
	// } else {
	// 	fmt.Println("Not elligible")
	// }

	// number := 30
	// if number > 30 {
	// 	fmt.Println("Number is greater than 30")
	// } else if number < 30 {
	// 	fmt.Println("Number is smaller than 30")
	// } else {
	// 	fmt.Println("Number is 30")
	// }

	// switch cases
	day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
		// fallthrough
	case 3:
		fmt.Println("Wednesday")
	case 4, 10, 20:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a valid input")

	}

	// number := "One"
	// switch number {
	// case "One":
	// 	fmt.Println(1)
	// case "Two":
	// 	fmt.Println(2)
	// case "Three":
	// 	fmt.Println(3)
	// }

	// No paranthesis around the condition
	//Starting bracket and condition should be in same line
	// if true {
	// 	fmt.Println("True")
	// }
	// if x := 10; x > 7 {
	// 	fmt.Println("True")
	// }

	// Fallthrough

	age := 30
	isAllowedToDrive := false

	if age > 20 && isAllowedToDrive {
		fmt.Println("Elligible for Driving license")
	} else {
		fmt.Println("Not elligible")
	}

	// break

}
