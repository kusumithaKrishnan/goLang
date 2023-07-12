package main

import "fmt"

func main() {
	fmt.Println("Maps in Go Lang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all lang", languages)

	delete(languages, "RB")
	fmt.Println("List of all lang", languages)

	//for in maps

	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}

}
