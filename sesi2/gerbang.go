package main

import "fmt"

func main() {
	gender := "male"
	age := 17

	var canWork bool

	// Penggunaan AND
	if gender == "male" && age >= 18 {
		canWork = true
	} else {
		canWork = false
	}

	fmt.Println("Nilai canWork pada notasi AND", canWork)

	// Penggunaan OR
	if gender == "male" || age >= 18 {
		canWork = true
	} else {
		canWork = false
	}

	fmt.Println("Nilai canWork pada notasi OR", canWork)

	// Penggunaan NOT
	if gender != "male" {
		canWork = true
	} else {
		canWork = false
	}

	fmt.Println("Nilai canWork pada notasi NOT", canWork)
}
