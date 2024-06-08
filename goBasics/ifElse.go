package main

import "fmt"

func main() {
	n := 7
	// if without else
	if 6%2 == 0 {
		fmt.Println(" 6 is divisible by 2")
	}

	if 3%2 == 0 || 8%2 == 0 {
		fmt.Println("either 3 or 8 are even")
	}

	if n%2 == 0 {

		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

}
