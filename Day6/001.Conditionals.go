package main

import "fmt"

func main() {

	fmt.Println("Conditionals")

	// decision making statement
	//  if else else if switch

	// if condition:
	x := 45
	if x > 30 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	// Top to bottom manner

	// if else-if else
	if x < 50 {
		fmt.Println("X is greater than 40")
	} else if x > 40 {
		fmt.Println("X is equal to 40")
	} else {
		fmt.Println("X is smaller than 40")
	}

	// Switch case
	day := 1
	switch day {
	case 0, 10, 100:
		fmt.Println("Sunday")
	case 1, 11, 111:
		fmt.Println("Monday")
		fallthrough
	case 15, 20, 30:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid Day number")
	}

	//fallthrough
	// temp := 100
	// switch {
	// case temp > 30:
	// 	fmt.Println()
	// }

	// Odd and Even Number

	number := 29
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Grade Calculator
	marks := 78
	if marks > 90 {
		fmt.Println("A")
	} else if marks > 80 {
		fmt.Println("B")
	} else if marks > 70 {
		fmt.Println("C")
	} else if marks > 50 {
		fmt.Println("D")
	} else {
		fmt.Println("Fail")
	}

	// Homework :
	// Convert the above logic in switch case
	// Write a program for Calcluator -- + , -, /, *
	// num1 num2
	// switch case  "+"
	// default Invalid Operation

}
