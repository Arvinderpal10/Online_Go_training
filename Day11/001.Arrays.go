package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")
	var myArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)
	// var myArray2 = [5]int{1,2,3}
	fmt.Println(len(myArray))

	//Short assignment
	// myArray3:=[5]int{1,2}
	// numbers := [...]int{1, 2, 3, 4, 5}
	// fmt.Println(len(numbers))

	// Arrays are zero index based
	// var myArray = [5]int{1, 2, 3, 4, 5}
	//                      0   1  2  3  4
	fmt.Println(myArray[3])

	// TO modify any element
	// myArray[3] = 145
	// fmt.Println(myArray[3])

	for i := 0; i < len(myArray); i++ {
		fmt.Println("Index:", i, "Values:", myArray[i])
	}
	for i, value := range myArray {
		fmt.Println("Index:", i, " Value:", value)

	}
	for _, value := range myArray {
		fmt.Println("Value:", value)
	}

	myArray3 := myArray
	// var myArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray3)

	matrixA := [3][3]int{
		{11, 22, 33},
		{12, 34, 56},
		{34, 78, 90},
	}

	matrixB := [3][3]int{
		{13, 2, 3},
		{5, 4, 6},
		{4, 8, 0},
	}

	result := [3][3]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
	for _, row := range result {
		fmt.Println(row)
	}

	number2 := [100]int{12, 34, 56, 78, 90, 13, 15, 14, 16, 46}
	evenCount := 0
	oddCount := 0
	evenNumber := []int{}
	oddNumber := []int{}
	for _, value := range number2 {
		if value%2 == 0 {
			evenNumber = append(evenNumber, value)
			evenCount++
		} else {
			oddNumber = append(oddNumber, value)
			oddCount++
		}
	}
	fmt.Println("Even Number:", evenNumber)
	fmt.Println("Odd Number:", oddNumber)
	fmt.Println("Even numbers :", evenCount, "Odd Numbers:", oddCount)

}
