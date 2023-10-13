package main

import (
"fmt"
"time"
)

func vals(a int, b int) (int, int){
	return a, b
}

func main() {
	start := time.Now()
	a, b := vals(5,9)
	
	fmt.Println(a)
	fmt.Println(b)
	
	_, c := vals(5,9)
	fmt.Println(c)
	d, _ := vals(5,9)
	fmt.Println(d)
	timeElapsed := time.Since(start)
	fmt.Println(timeElapsed)
}
