package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 2)

	wg.Add(2)

	go runMain1(ch)
	go runMain2(ch)

	wg.Wait()

	fmt.Println("Ending Main Function......")
}

func runMain1(mk chan int) {
	defer wg.Done()
	fmt.Println("Adding Value to Channel...")
	mk <- 10
}

func runMain2(mk chan int) {
	defer wg.Done()
	fmt.Println("Receiving Value From Channel...")
	fmt.Println(<-mk)
}

// Buffered Channel does not garanteed of massege delivery.

// it have length as compared to Unbuffered Channel..
