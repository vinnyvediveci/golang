package main

import "fmt"

func main() {
	var a = "initial string"
	fmt.Println(a)

	// var declares 1 or more variables.
	var b, c int = 1, 2
    	fmt.Println(b, c)
	
	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)
	
	// Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.
	f := "apple, declared and initialized"
	fmt.Println(f)

}
