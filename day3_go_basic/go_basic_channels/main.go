package main

import (
	"fmt"
	"math/rand"
	"time"
)
var MAX_OBJECT_PRICE float32 = 7
var MAX_TOFU_PRICE float32 = 5
func main() {
	var objectChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"dmart.com", "bigbasket.com", "amazon.com", "flipkart.com"}
	for i:= range websites{
		go checkObjectPrices(websites[i], objectChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}

	sendMessage(objectChannel, tofuChannel)
}

func checkTofuPrices(s string, tofuChannel chan string) {
	for {
		time.Sleep(1 * time.Second)
		var tofuPrice = rand.Float32()*20
		fmt.Println("Price of tofu",tofuPrice,"on",s)
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- s
			break
		}
	}
}



func checkObjectPrices(website string, objectChannel chan string) {
	for {
		time.Sleep(1 * time.Second)
		var objectPrice = rand.Float32()*20
		fmt.Println("Price of object",objectPrice,"on",website)
		if objectPrice <= MAX_OBJECT_PRICE {
			objectChannel <- website
			break
		}
	}
}

func sendMessage(objectChannel chan string, tofuChannel chan string) {
	select{
		case website := <-objectChannel:
			fmt.Printf("\nText Sent: Found deal on Object at %v.", website)
		case website := <-tofuChannel:
			fmt.Printf("\nText Sent: Found deal on Tofu at %v.", website)
	}
}