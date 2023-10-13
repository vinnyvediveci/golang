package main

import "fmt"

func vals(a int, b int) (int, int){
	return a, b
}

func main() {
	a, b := vals(5,9)
	
	fmt.Println(a)
	fmt.Println(b)
	
	_, c := vals(5,9)
	fmt.Println(c)
	d, _ := vals(5,9)
	fmt.Println(d)
}
