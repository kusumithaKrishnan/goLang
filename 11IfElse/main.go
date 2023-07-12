package main

import "fmt"

func main() {
	fmt.Println("If Else")

	loginCount := 23
	var result string

	if loginCount > 10 {
		result = "Greater than 10"
	} else if loginCount < 10 {
		result = "Less than 10"
	} else {
		result = "exactly 10"
	}

	fmt.Println("Results are, ", result)

	// we can directly assign on fly

	if 9%2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is odd")
	}

	if num := 3; num > 10 {
		fmt.Println("num is greater than 10")
	} else {
		fmt.Println("Num is less than 10")
	}
}
