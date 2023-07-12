package main

import "fmt"

func main() {
	fmt.Println("Welocome to Array")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Orange"
	fmt.Println("FruitList is ", fruitList)
	fmt.Println("FruitList is ", len(fruitList))

	var vegList = [3]string{"Beans", "Carrot", "Pumkin"}
	fmt.Println("Veg List", len(vegList))

}
