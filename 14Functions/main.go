package main

import "fmt"

func main() {
	fmt.Println("This is the main function")
	greeter()
	result := adder(3, 2)
	fmt.Println("The Result is :", result)
	proAdd := proAdder(1, 2, 3, 4)
	fmt.Println("The sum is :", proAdd)
}

func greeter() {
	fmt.Println("This is greeter function")
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
