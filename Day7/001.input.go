//  Command line inputs

// INput -- fmt.Scan() --- Single line inputs , Space Separated

package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	var name string
// 	var age int
// 	fmt.Print("Enter your name and age: ") // Shujath 27
// 	fmt.Scan(&name, &age)
// 	fmt.Println(name, age)

// }

// fmt.Scanln() --- It stops with new line
// func main() {
// 	var name string
// 	var age int
// 	fmt.Print("Enter your name and age:")
// 	fmt.Scanln(&name, &age)
// 	fmt.Print(name, " ", age)
// }

// fmt.Scanf() --- Formatting -- %s , %v,
// func main() {
// 	var name string
// 	var age int
// 	fmt.Print("Enter your name and age :")
// 	fmt.Scanf("%s %d", &name, &age)
// 	// fmt.Print(name, " ", age)
// 	fmt.Printf("%s  %d ", name, age)

// }

// bufio.Reader() -- Fullline , put a delimeter
// 4096 bytes
// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter your details :")
// 	sentence, _ := reader.ReadString('\n')
// 	fmt.Print(sentence)

// }

// bufio.Scanner() --  Read files line by line
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	fmt.Println("Enter you input:")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	fmt.Println("Result:")
	for _, l := range lines {
		fmt.Print(l)
	}

}
