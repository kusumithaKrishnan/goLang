package main

import "fmt"

func main() {
	fmt.Println("Methods")

	kusumitha := User{"Kusum", "KK@dev.go", true, 28}

	fmt.Println("Name Obj is ", kusumitha)
	fmt.Printf("Details Are %+v\n", kusumitha)

	// %+v to get all the data
	// $v to get single data

	fmt.Printf("Name is %v", kusumitha.Name)

	kusumitha.GetStatus()
	kusumitha.GetEmail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("The status is", u.Status)
}

func (u User) GetEmail() {
	u.Email = "Sakthi@dev.com"
	fmt.Println("The New email is", u.Email)
}
