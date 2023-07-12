package main

import "fmt"

func main() {
	fmt.Println(("Welcome to Loops"))

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for key, value := range days {
		fmt.Printf("The Key is %v and value is %v\n", key, value)
	}

	// we have while , go while, break, goto -- checkout in youtube
}
