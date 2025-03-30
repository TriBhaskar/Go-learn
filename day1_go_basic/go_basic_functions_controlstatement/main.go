package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("John Doe")
	var res, remainder, err = intDivision(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	}else if remainder == 0 {
		fmt.Println("result of intDivision",res)
	}else {
		fmt.Println("remainder of intDivision",remainder)
	}

	switch remainder{
	case 0:
		fmt.Println("remainder is zero")
	case 1,2:
		fmt.Println("remainder is 1 or 2")
	default:
		fmt.Println("remainder is something else")
	}
	
}

func printMe(printValue string){
	fmt.Printf("Hello, %s here!", printValue)
	fmt.Println("")
}

func intDivision(numerator int, denominator int ) (int , int, error) {
	
	var err error
	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0 , 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}