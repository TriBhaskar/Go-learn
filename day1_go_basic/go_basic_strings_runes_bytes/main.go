package main

import (
	"fmt"
	"strings"
)

func main() {

	// String in go
	var myString = []rune("Bhèskar")
	fmt.Println(myString)
	var indexed = myString[2] 
	fmt.Printf("%v, %T\n", indexed, indexed) // 66, uint8

	for i,v := range myString{
		fmt.Println("index:", i, "value:", v) // index: 0 value: 66
	}

	// Rune
	var myRune rune = 'B'
	fmt.Println("myrune",myRune) // 66

	var strSlice = []string{"b","h", "è", "s", "k", "a", "r"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Println("catStr", catStr) // catStr Bhèskar

	// we cannot update the string as it is immutable

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	var catStr2 = strBuilder.String()
	fmt.Println("catStr2", catStr2) // catStr2 Bhèskar


}