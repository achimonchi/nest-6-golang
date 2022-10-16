package main

import "fmt"

func main() {
	var num1, num2 float32 = 10, 20
	var age float32 = 2

	text := "Hello" + " World"
	i := 0
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)

	age += num1

	fmt.Println(age)
	fmt.Println(num1+num2, num1-age, num2*age, num1/age, text)
}
