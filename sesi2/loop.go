package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("While ...")
	counter := 0
	for counter < 5 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("do while")
	for {
		fmt.Println(counter)
		counter--
		if counter < 0 {
			break
		}
	}

	fruits := []string{"Grape", "Banana", "Apple"}

	for key, val := range fruits {
		fmt.Println("Key =>", key)
		fmt.Println("Val =>", val)
	}
	fmt.Println("done...")
}
