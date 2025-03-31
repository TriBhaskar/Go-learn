package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
	owner
}

func (g gasEngine) milesLeft() uint8 {
	return g.gallons * g.mpg  
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(g engine, miles uint8) {
	if miles <= g.milesLeft() {
		fmt.Println("You can make it!")
	}else{
		fmt.Println("You cannot make it!")
	}
}

// interface{} is a type that can hold any value
type engine interface {
	milesLeft() uint8
}

type owner struct{
	name string
}

func main() {

	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 10, owner: owner{name: "John Doe"}}
	myEngine.mpg = 30 // This will not work, as gasEngine is a struct type, not a variable
	fmt.Println("Gas Engine:", myEngine.gallons, "gallons:" , myEngine.mpg, "mpg:", myEngine.name)
	fmt.Println("Miles driven:", myEngine.milesLeft())
	canMakeIt(myEngine, 200)
	//anonymous struct
	var myEngine2 = struct {
		mpg uint8
		gallons uint8
	}{mpg: 25, gallons: 10}

	fmt.Println("Anonymous Engine:", myEngine2.gallons, "gallons:" , myEngine2.mpg, "mpg:")

	var myElectricEngine electricEngine = electricEngine{mpkwh: 4, kwh: 10}
	fmt.Println("Electric Engine:", myElectricEngine.kwh, "kwh:" , myElectricEngine.mpkwh, "mpkwh:")
	fmt.Println("Miles driven:", myElectricEngine.milesLeft())
	canMakeIt(myElectricEngine, 10)
}