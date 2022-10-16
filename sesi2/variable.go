package main

import "fmt"

var name string = "Hello"

func main() {
	// var age int32
	// age = 20

	address, gender, age := "Jakarta", "Male", 10
	var isOdd bool = false

	text := `
		Hello world
		Ini adalah variable text
	`

	// if address == "" {
	// 	text := ""
	// }

	// fmt.Println(text)
	sayHello()
	fmt.Println(text)
	fmt.Println(name, age, address, gender, isOdd)
}

func sayHello() {
	fmt.Println("Hello", name)
}
