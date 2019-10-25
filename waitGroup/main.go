package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Main Function Start Running........")

	wg.Add(2)

	go goRun1()
	go goRun2()
	wg.Wait()
	fmt.Println("Main Function Ended Running........")
}

func goRun1() {
	defer wg.Done()
	fmt.Println("Starting goRun1() function")
	for i := 0; i < 5; i++ {
		fmt.Println("goRun1 Running", i)
	}
	fmt.Println("goRun1() Executed Completely and Now Ended")
}

func goRun2() {
	defer wg.Done()
	fmt.Println("Starting goRun() function....")
	for i := 0; i < 5; i++ {
		fmt.Println("goRun2 Running", i)
	}
	fmt.Println("goRun2() Executed Completely and Now Ended")
}
