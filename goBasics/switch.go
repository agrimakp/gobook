package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println(" write", i, " as ")
	// basic switch
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's Working day")
	}

}
