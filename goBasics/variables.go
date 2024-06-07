package main

import "fmt"

var global = "global"

func main() {

	fmt.Println(global)

	var a = "initial"
	fmt.Println(a)
	var g = 3
	fmt.Println(g)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)
	// 0
	var h string
	fmt.Println(h)
	// empty
	var i bool
	fmt.Println(i)
	// false

	//  shorthand for declaring and initializing a variable
	f := "hai"
	fmt.Println(f)
}
