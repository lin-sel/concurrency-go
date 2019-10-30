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

	go runRoutine2(mk)

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
