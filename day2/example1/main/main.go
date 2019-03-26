package main

import (
	"fmt"
	"math"
)

func test(n int) int {
	tmp := int(math.Sqrt(float64(n)))
	return tmp
}

func main() {
	fmt.Println(test(8))
}
