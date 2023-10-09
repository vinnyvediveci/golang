// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages) https://en.wikipedia.org/wiki/Associative_array

package main

import (
	"fmt"
//	"maps"
)

func main() {

	m := make(map[string]int)
	
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)


}
