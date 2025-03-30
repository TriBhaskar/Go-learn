package main

import "fmt"

func main() {

	// array declaration and initialization
	var intArr [3]int32
	intArr[0] = 1
	intArr[1] = 2
	fmt.Println("array values arr[0]:", intArr[0])
	fmt.Println("array values of 1 & 2nd index:", intArr[1:3])

	//print memory address of array
	fmt.Printf("array address: %p\n", &intArr[0])
	fmt.Printf("array address: %p\n", &intArr[1])
	fmt.Printf("array address: %p\n", &intArr[2])

	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println("Array2 ", intArr2)

	intArr3 := [3]int32{4, 5, 6}
	fmt.Println("Array3 ", intArr3)

	// slice declaration and initialization
	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Printf("the length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 9)
	fmt.Printf("the length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 =[]int32{4, 5, 6}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 5, 10)
	fmt.Println("intSlice3", intSlice3)

	// map declaration and initialization

	var myMap map[string]int32 = make(map[string]int32)
	fmt.Println("myMap", myMap)

	var myMap2 map[string]int32 = map[string]int32{"a": 1, "b": 2}
	fmt.Println("myMap2", myMap2)
	fmt.Println("myMap2 value of a", myMap2["a"])
	fmt.Println("myMap2 value of b", myMap2["b"])
	fmt.Println("myMap2 value of c", myMap2["c"])
	var value, ok = myMap2["c"]
	fmt.Println("myMap2 value of c", value, ok)

	delete(myMap2, "a")
	fmt.Println("myMap2", myMap2)

	var myMap3 = map[string]int32{"john":30, "doe": 40, "jane": 50}

	// looping through map
	fmt.Println("looping through map")
	for key,value := range myMap3 {
		fmt.Println("name:", key, "age:", value)
	}

	// looping through array
	fmt.Println("looping through array")
	for index, value := range intArr3 {
		fmt.Println("index:", index, "value:", value)
	}

	// while loop in go
	fmt.Println("while loop")
	var i int =0
	for i<10{
		fmt.Println("i:", i)
		i++
	}

	// for loop in go
	fmt.Println("for loop")
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}



}