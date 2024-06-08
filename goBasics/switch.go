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
	
	// switch with an expression
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's Working day")
	}

	//  switch without an expression

	t := time.Now()

	switch {
	  case t.Hour() < 12:
		  fmt.Println("It's before noon")
          default: 
	        fmt.Println("")
	}
}
