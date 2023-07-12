package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")
	// When you pass on a variable to a function or class, sometimes the copy of the variable is passed on and not the actual value.
	// so to avoid this we use pointer. we pass memory address of the vaiable , so that when passed its teh actual value from the memory.

	// Direct address of a memory address == pointer definition

	var one *int // pointer which stores the value as integer

	fmt.Println("Value of one", one)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("value of actual pointer is ", ptr)

}
