package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"apple", "tomato", "peach"}

	fmt.Printf("Type of fruitList %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("FruitList is", fruitList)

	// fruitList = append(fruitList[1:])   // start value is added and rest of the data will come as it is
	// fmt.Println("FruitList Trimed", fruitList)

	// fruitList = append(fruitList[1:3]) //start and end value is given
	// fmt.Println("FruitList Trimed", fruitList)

	fruitList = append(fruitList[:3]) // only end value is given. IN this case start value is considered as 0
	fmt.Println("FruitList Trimed", fruitList)

	// make()

	highScores := make([]int, 4)

	highScores[0] = 111
	highScores[1] = 999
	highScores[2] = 222
	highScores[3] = 444
	// highScores[5] = 211

	highScores = append(highScores, 888, 777, 666)

	sort.Ints(highScores)

	fmt.Println("High Scroes", highScores)
	fmt.Println("is sorted", sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"react", "Go", "ReactNative", "Next", "angular", "java"}

	fmt.Println("Courses", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses", courses)
}
