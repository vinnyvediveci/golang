package main

import "fmt"


func main() {

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
		} else if num < 10 {
			fmt.Println(num, "is one digit")
		} else {
			fmt.Println(num, "has several digits")
		}
		 
}

