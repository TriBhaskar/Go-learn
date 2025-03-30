package main

import "fmt"
func main() {
	// This is a placeholder for the main function.
	// You can add your code here to run the application.
	fmt.Println("Hello, World!")
	// For example, you can call other functions or execute logic here.

	var intVar int
	intVar = 42
	fmt.Println("Integer variable:", intVar)

	var floatVar float64 = 123456.14
	fmt.Println("Float variable:", floatVar)

	var floatVar2 float32 = 123456.14
	var intVar2 int = 123456
	var result float32 = floatVar2 + float32(intVar2)
	fmt.Println("Result of float32 + int:", result)

	//now division of int 

	var intVar3 int = 10
	var intVar4 int = 3
	fmt.Println("Integer division:", intVar3/intVar4) // This will perform integer division
	fmt.Println("Integer modulus:", intVar3%intVar4) // This will give the remainder of the division

	var myString string = "Bhaskar Ghosh"
	fmt.Println("String variable:", myString)
	//count characters in string function
	fmt.Println("Length of string:", len(myString))

	var myRune rune = 'A'
	fmt.Println("Rune variable:", myRune)
	fmt.Println("Rune as string:", string(myRune))

	var myBool bool = true
	fmt.Println("Boolean variable:", myBool)

	var myVar = "text"
	fmt.Println("Variable with inferred type:", myVar)

	myNewVar := "new text"
	fmt.Println("Variable with short declaration:", myNewVar)

	var v1, v2, v3 int = 1, 2, 3
	fmt.Println("Multiple variable declaration:", v1, v2, v3)

	const myConst = "constant value"
	fmt.Println("Constant variable:", myConst)

	const pi float32 = 3.14
	fmt.Println("Constant pi:", pi)
}