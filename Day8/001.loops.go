package main

import "fmt"

//  Initilization -- One time

//  increement -- n+1

// Conditon check -- n+1
func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// While
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// num := []int{10, 12, 63, 64}
	// for index, value := range num {
	// 	fmt.Printf("Index: %d , value: %d \n", index, value)

	// }

	// num := []int{10, 12, 63, 64}
	// for _, value := range num {
	// 	fmt.Printf("Value: %d \n", value)

	// }

	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Print(" ", i*j)
	// 	}
	// 	fmt.Println()
	// }

	// for i := 0; i < 5; i++ {
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Print(" ", i)
	// }

	for i := 0; i < 10; i++ {
		if i < 5 {
			continue
		}
		fmt.Print(" ", i)
	}
}
