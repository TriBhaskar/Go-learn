package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg float32
}

type electricEngine struct {
	kwh float32
	mpkwh float32
}

type car [T gasEngine | electricEngine] struct {
	carMake string
	carModel string
	engine T
}

func main() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of int slice:", sumSlice(intSlice))

	var floatSlice = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Sum of float64 slice:", sumSlice(floatSlice))

	var intVar = []int{}
	fmt.Println("Is int slice empty?", isEmpty(intVar))
	var floatVar = []float64{2.2, 3.3, 4.4}
	fmt.Println("Is float64 slice empty?", isEmpty(floatVar))


	var gasCar = car[gasEngine]{
		carMake: "Toyota",
		carModel: "Camry",
		engine: gasEngine{
			gallons: 15.0,
			mpg: 30.0,
		},
	}

	var electricCar = car[electricEngine]{
		carMake: "Tesla",
		carModel: "Model S",
		engine: electricEngine{
			kwh: 100.0,
			mpkwh: 4.0,
		},
	}

	fmt.Println("Gas Car:", gasCar)
	fmt.Println("Electric Car:", electricCar)

}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}