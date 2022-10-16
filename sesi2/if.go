package main

import "fmt"

func main() {
	isMale := true
	age := 16

	var canWork bool
	if isMale == true {
		if age >= 17 {
			canWork = true
		}
		fmt.Println("Laki laki")
	} else if isMale == false {
		if age >= 20 {
			canWork = true
		}
		fmt.Println("Perempuan")
	}

	fmt.Println(canWork)
}
