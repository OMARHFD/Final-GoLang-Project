package main

import (
	"fmt"
)

func main() {
	backends := []bool{false, false, false, false, true}
	fmt.Println(test(0, backends))

}

func test(c int, backends []bool) bool {
	counter := 0
	initialValue := c
	for !backends[c] && counter < (len(backends)) {
		c++
		counter++
		c = c % (len(backends))

	}

	if backends[c] {
		fmt.Println(c)
		return backends[c]
	} else {
		c = initialValue
		fmt.Println("error handling")
		return false
	}

}
