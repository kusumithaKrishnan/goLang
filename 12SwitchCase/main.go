package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and move 2 ")
	case 3:
		fmt.Println("Dice value is 3 and move 3")
	case 4:
		fmt.Println("Dice value is 4 and move 4")
	case 5:
		fmt.Println("Dice value is 5 and move 5")
	case 6:
		fmt.Println("Dice value is 6 and move 6")
	default:
		fmt.Println("Something wrong ")

	}

}
