package main

import "fmt"

func main() {
	fmt.Println("Operators in GOlang")

	//  Logical - And Or Not
	// && - True && True -- True  == AND
	// || - True || True , True || False
	// NOT
	a := true
	b := true
	fmt.Println(a && b) // True
	c := false
	fmt.Println(a && c) // False

	fmt.Println(a || b) // True
	fmt.Println(a || c) // True

	fmt.Println(!a)

	// We use logical operators to combine two or more conditions

	// Comparision :
	// less than . grater than , less rthan equal to , greater than equal to , ==

	// True false
	fmt.Println(5 == 3)
	fmt.Println(5 > 3)
	fmt.Println(6 < 7)
	fmt.Println(3 >= 4)
	fmt.Println(5 <= 7)

	// Bitwise Operator -- & , | , ^ ,<< , >>
	//  5 -- 101    0 0 1
	// 3 -   011
	// 1 & 1 = 1
	// 1 & 0 = 0
	// 0 & 0 = 0

	// 1    0    1
	// 0    1    1
	// 0      0    1

	fmt.Println(5 & 3) //1

	// 1 | 1 = 1
	// 1 | 0 = 1
	// 0 | 0 = 0

	// 1 1  1
	fmt.Println(5 | 3)

	// 5 - 1 0 1

	// left shift-- >  right shift

	//  1 0 1 0 0 -- 5 << 2
	// 10    -- Left Shift == 5 * 2 ^2 = 20
	fmt.Println(5 << 2)

	//  pointers -- & -- We are storing the address
	// *  -- > We are fetching the value at address
	x := 59
	p := &x

	// x = 59 ,
	// p = &x
	x = 40
	// y := 56

	fmt.Println(x, p, *p)

	// <-  send/ recieve value from the channel

}
