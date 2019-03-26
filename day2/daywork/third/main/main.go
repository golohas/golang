package main

import "fmt"

//var sum int

func main() {
	s := factorial(10)
	fmt.Println("factorial sum:")
	fmt.Println(s)

}
func factorial(n int) int {
	var sum int
	var m int = 1
	for i := 1; i <= n; i++ {
		m = m * i
		sum = sum + m
	}
	return sum
}
