package main // The main package where the entry point of the program is defined

import "fmt" // Importing the 'fmt' package for input and output operations

func main() { // The main function where the execution starts

	// fmt.Println() is used to print a new line after the string/values are displayed
	// Here, we're using Print() which does not add a new line, while fmt.Println() adds a new line after printing.

	fmt.Print("Hello World") // Prints "Hello World" without a new line at the end

	fmt.Println("Hello") // Prints "Hello" followed by a new line

	fmt.Println(30 + 10) // Prints the result of 30 + 10, which is 40

	fmt.Println("30", "10") // Prints the strings "30" and "10", separated by a space

	// ASCII values
	// H - 48 - Binary
	// ASCII is used to represent characters as numbers (H = 48 in ASCII)

	// Variable Declaration
	// A variable is a container that holds a value of a particular type.

	// 1. Declaration (without initialization)
	var x int      // Declaring 'x' as an integer (default value will be 0)
	fmt.Println(x) // Prints: 0 (default value for int)

	// 2. Initialization (declaration + initialization)
	var y int = 100 // Declaring 'y' and initializing it with 100
	fmt.Println(y)  // Prints: 100

	// 3. Type Inference (Go automatically infers the type)
	z := 200       // Go infers 'z' as an int because 200 is an integer
	fmt.Println(z) // Prints: 200

	// Naming Conventions
	// 1. A variable name can't start with a digit (e.g., "3Class" is invalid).
	// 2. Special characters are not allowed except for underscores '_'.
	// Examples of valid names: age2, name3Class, _name.
	// Invalid: 3Class, @Name.

	// You cannot use keywords as variable names. '_var' is a valid name, but 'var' is a keyword.

	_var := 1000      // Declaring and initializing '_var' with 1000 using shorthand syntax
	fmt.Println(_var) // Prints the value of '_var', which is 1000

	// Descriptive variable names are important for code readability.
	// For example, naming a variable 'age' is more meaningful than just 'x'.

	// Case sensitivity: Go is case-sensitive, meaning "Age" and "age" are considered different variables.

	// Print statements are case-sensitive as well. 'Print()' and 'print()' are different.
	//fmt.Println(age)   // Prints the value of 'age' (24)
	//fmt.Println(marks) // Prints the value of 'marks' (100)
	//fmt.Println(u)     // Prints the value of 'u' (100)

	// Unsed variables will throw an error.

	// Naming conventions, case sensitivity, and shorthand declarations are useful in Go.

}
