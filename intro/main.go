package main

import "fmt"

/*
Goroutines are functions or methods that run concurrently with other functions or methods.
Goroutines can be thought of as light weight threads.
The cost of creating a Goroutine is tiny when compared to a thread.
Hence its common for Go applications to have thousands of Goroutines running concurrently.
*/
func main() {
	fmt.Println("Starting Main Function Execution.......")
	// Start func execution on other core.
	go goRun()
	fmt.Println("Ending Main Function Execution......")
}

func goRun() {
	fmt.Println("Start GoRun function Execution.....")
	fmt.Println("Ending GoRun function Execution.....")
}

//Output:
/*
Starting Main Function Execution.......
Ending Main Function Execution......
*/

// Explaination:
// In this whole Program there is two go routine. one is for main function and other for goRun function.
// main function go routine start execution then after 1st statement print then it start goRun() function on the other core. After calling goRun() function Execution control
// come back to main function.
// then it execute last statement of the main function. after that main function terminated.
// Here Problem is Main function don't care about goRun() function. whether it executed or not. main function will not wait for goRun function execute completely.
// when main functions go routine terminated then no other go routine will executing.
