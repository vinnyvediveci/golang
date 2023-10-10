// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages) https://en.wikipedia.org/wiki/Associative_array

package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)
	
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"] 
	fmt.Println("v3:", v3)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("del1:", m)
	
	clear(m)
	fmt.Println("clear:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map1:", n1)

	n2 := map[string]int{"foo": 1, "bar":2}
	fmt.Println("map2:", n2)

	if maps.Equal(n1, n2) {
		fmt.Println("maps n1 and n2 are equal")
	} 
}
