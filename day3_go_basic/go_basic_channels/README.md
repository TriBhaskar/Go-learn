# Go Channels Example - Price Checker

This project demonstrates the usage of Go channels in a price checking simulation. The program checks prices of items across different e-commerce websites concurrently.

## Key Channel Concepts Demonstrated

1. **Channel Creation**

   ```go
   var objectChannel = make(chan string)
   var tofuChannel = make(chan string)
   ```

   - Channels are created using the `make` function
   - In this example, we create channels that transmit string values

2. **Goroutines with Channels**

   ```go
   for i:= range websites {
       go checkObjectPrices(websites[i], objectChannel)
       go checkTofuPrices(websites[i], tofuChannel)
   }
   ```

   - Multiple goroutines run concurrently
   - Each goroutine monitors prices on different websites
   - Channels enable communication between goroutines

3. **Channel Send Operation**

   ```go
   if objectPrice <= MAX_OBJECT_PRICE {
       objectChannel <- website  // Send website to channel
       break
   }
   ```

   - Uses arrow operator (`<-`) to send data into channel
   - Sender blocks until receiver is ready

4. **Channel Select Statement**
   ```go
   select {
       case website := <-objectChannel:
           fmt.Printf("\nText Sent: Found deal on Object at %v.", website)
       case website := <-tofuChannel:
           fmt.Printf("\nText Sent: Found deal on Tofu at %v.", website)
   }
   ```
   - `select` allows listening on multiple channels
   - First channel with available data is processed
   - Non-blocking way to handle multiple channels

## Program Flow

1. Program creates two channels for monitoring object and tofu prices
2. Launches goroutines to check prices on multiple websites
3. Each goroutine continuously checks prices until finding one below threshold
4. When a good price is found, website URL is sent through channel
5. Main program uses `select` to receive and process the first good deal found

## Usage

Run the program:

```bash
go run main.go
```

The program will continuously print prices from different websites and stop when it finds the first deal (either object <= $7 or tofu <= $5).
