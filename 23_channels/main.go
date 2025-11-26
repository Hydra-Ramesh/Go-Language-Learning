package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChain chan int) {
	for num := range numChain {
		fmt.Println("Received number:", num)
		time.Sleep(500 * time.Millisecond)
	}
}

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing...")
	time.Sleep(300 * time.Millisecond)
}

func emailSender(emailChain <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChain {
		fmt.Println("Sending email to:", email)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Example 1: processNum
	fmt.Println("Example 1: processNum")
	numChain := make(chan int)
	go processNum(numChain)
	for i := 0; i < 5; i++ {
		numChain <- rand.Intn(100)
	}
	close(numChain)
	// give the goroutine time to finish processing the closed channel
	time.Sleep(1 * time.Second)

	// Example 2: sum
	fmt.Println("\nExample 2: sum")
	result := make(chan int)
	go sum(result, 3, 5)
	fmt.Println("Sum result:", <-result)

	// Example 3: task
	fmt.Println("\nExample 3: task")
	done := make(chan bool)
	go task(done)
	<-done

	// Example 4: emailSender
	fmt.Println("\nExample 4: emailSender")
	emailChain := make(chan string, 100)
	doneEmail := make(chan bool)
	go emailSender(emailChain, doneEmail)
	for i := 0; i < 10; i++ {
		emailChain <- fmt.Sprintf("user%d@example.com", i)
	}
	emailChain <- "ramesh@gmail.com"
	emailChain <- "rakesh@gmail.com"
	close(emailChain)
	<-doneEmail
	fmt.Println("All emails processed.")

	// Example 5: select with two channels
	fmt.Println("\nExample 5: select")
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() { chan1 <- 42 }()
	go func() { chan2 <- "Hello, Channels!" }()
	for i := 0; i < 2; i++ {
		select {
		case num := <-chan1:
			fmt.Println("Received from chan1:", num)
		case msg := <-chan2:
			fmt.Println("Received from chan2:", msg)
		}
	}
}
