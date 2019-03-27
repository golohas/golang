package main

import (
	"fmt"
	"strings"
)

func main() {
	// TODO strings replace function
	var str, str1 string
	str = "abbacba"
	str1 = " hello world abc "

	trimStr := strings.Trim(str, "ab")
	fmt.Println(trimStr)
	count := strings.Count(str, "ab")
	fmt.Println(count)
	fmt.Println(strings.Replace(str1, "world", "you", 1))
	fmt.Println(strings.Repeat(str, 3))
	fmt.Println(strings.TrimSpace(str1))
	fmt.Println(strings.Trim(str1, " "))
}
