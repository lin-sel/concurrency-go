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

// In this function, function have only write permission.
func runMain1(mk chan<- int) {
	defer wg.Done()
	fmt.Println("Adding Value to Channel...")
	mk <- 10
}

// In this function, function have only read permission.
func runMain2(mk <-chan int) {
	defer wg.Done()
	fmt.Println("Receiving Value From Channel...")
	fmt.Println(<-mk)
}
