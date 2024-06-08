package main

import (
	"fmt"
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
}
