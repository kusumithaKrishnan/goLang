package main

import "fmt"

func main() {
	fmt.Println("Structs")

	kusumitha := User{"Kusum", "KK@dev.go", true, 28}

	fmt.Println("Name Obj is ", kusumitha)
	fmt.Printf("Details Are %+v\n", kusumitha)

	// %+v to get all the data
	// $v to get single data

	fmt.Printf("Name is %v", kusumitha.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
