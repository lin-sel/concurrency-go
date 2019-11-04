package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// Declaring Channel
	mk := make(chan int)

	wg.Add(2)

	// Run goroutine to set data on channel.
	go runRoutine2(mk)

	// Run Goroutine to receive data from channel.
	go runRoutine1(mk)

	wg.Wait()

	fmt.Println("Ending All Goroutine......")
}

func runRoutine2(mk chan int) {
	defer wg.Done()

	fmt.Println("Adding Value to Channel")
	mk <- 200
}

func runRoutine1(mk chan int) {
	defer wg.Done()
	fmt.Println("Channel Value", <-mk)
}

// In This Way Different Goroutine Communicate To each Other.


// Benefits Of Unbuffered Channel.

// It Guaranteed for Massage Delivery.