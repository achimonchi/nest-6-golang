package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	day := t.Weekday()

	isWin := true

	if isWin {

	}

	switch {
	case t.Hour() > 12 && (day == time.Saturday || day == time.Sunday):
		fmt.Println("Weekend siang")
	case t.Hour() < 12 && (day == time.Saturday || day == time.Sunday):
		fmt.Println("Weekend pagi")
	default:
		fmt.Println("Weekdays pagi / siang")
	}
}
