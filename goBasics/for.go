package main

import "fmt"

func main() {
	i := 1

	count := 3

	for i <= count {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < count; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	// for without a condition will loop repeatedly until you break out of the loop
	for {
		fmt.Println("loop repeatedly until you break")
		break
	}

}
