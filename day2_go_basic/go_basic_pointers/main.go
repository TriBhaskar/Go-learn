package main

import (
	"fmt"
)

func main(){

	//pointers
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)
	*p = 42
	fmt.Printf("\nThe value p points to is: %v", *p)
	p = &i
	fmt.Printf("\nThe value p points to is: %v", *p)
	*p = 100
	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)

	// *pointer means dereference the pointer, so you can get the value it points to
	// &pointer means get the address of the variable, so you can assign it to a pointer

	//slces	

	var slice = []int32{1, 2, 3, 4, 5}
	var slicyCopy = slice;
	slicyCopy[2] = 100
	fmt.Println("\n",slice)
	fmt.Println(slicyCopy)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n The memory location of the thing1 array is %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\n The value of thing1 is %v", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\n The memory location of the thing2 array is %p", &thing2)
	for i := range thing2{
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
