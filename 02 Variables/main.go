package main

import "fmt"

const LoginToken string = "Kusumitha"

func main() {
	var userName string = "kusum"
	fmt.Println(userName)
	fmt.Printf("Show the type of userName:  %T \n", userName)

	var isLoggedInt bool = true
	fmt.Println(isLoggedInt)
	fmt.Printf("Show the type of isLoggedInt:  %T \n", isLoggedInt)

	var val uint8 = 255
	fmt.Println(val)
	fmt.Printf("Show the type of val:  %T \n", val)

	var floatVal float32 = 255.3355554878
	fmt.Println(floatVal)
	fmt.Printf("Show the type of floatVal:  %T \n", floatVal)

	var floatVal64 float64 = 255.3355554878888
	fmt.Println(floatVal64)
	fmt.Printf("Show the type of floatVal64:  %T \n", floatVal64)

	// default values

	var defaultVal int
	fmt.Println(defaultVal)
	fmt.Printf("Show the type of defaultVal:  %T \n", defaultVal)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)

	// no var style

	noOfUsers := 300000 // := is walrun operator
	fmt.Println(noOfUsers)

	// using a public variable

	fmt.Println(LoginToken)
	fmt.Printf("Show the type of defaultVal:  %T \n", LoginToken)

}
