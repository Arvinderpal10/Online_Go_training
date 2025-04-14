package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("variables")

	// Tokens -- Indentifiers , keywords , Constnats , Punctuations , Operators

	//  Identifiers -- name of variable / constant / function
	// var x int = 40
	// fmt.Println(x)

	// Literals -- Values of the variable
	//  integral literals - 40
	//   float literals - 40.50
	//  string literals - "Hello World"

	// var var int = 200
	// fmt.Println(var)

	// var var3 int = 400

	//  Keywords - for if else var func

	// Semicolons are rarely used , Line seperators ; ,, for

	// data Types - int , float , string , boolean , rune
	// var x int = 100     // var is a keyword , x is an indetifire,  int data type , 100 -- literal

	var x = 100.30 // type inference
	fmt.Println(x)

	// Short hand declaration :
	y := 300

	// Type casting // type comnversion
	var z int = int(x)
	fmt.Println(y)
	fmt.Println(z)

	// Integer - int8 , int16 , int32 , int64 -- int
	// int8 - -128  ----- 127

	// uint - unsigned -- Only positive   -- uint8 -- 0 - 255

	//  float32 , float64  -- float64
	var u float64 = 300.5607
	fmt.Println(u)

	// Boolean -- true , false
	// fmt.Println(int(true))
	// Default --

	var v int // 0
	fmt.Println(v)
	var t float64 // 0.0
	fmt.Println(t)
	var b bool // false
	fmt.Println(b)
	var s string // ""
	fmt.Println(s)

	//  true - any number except 0
	// false - 0 , empty string
	fmt.Println(math.MaxInt8, math.MinInt8)

	// int8 -- 127
	// var r int16 = 137
	var num uint8 = 255 // Max value for uint8     0-255 -- 0-10 = 10
	fmt.Println("Before Overflow:", num)
	num = num + 10                      // Causes overflow
	fmt.Println("After Overflow:", num) // Wraps to 0

}

// go run filename.extension
