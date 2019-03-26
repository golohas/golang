package main

import "fmt"

func main() {
	fmt.Println("flower shu: ")
	for i := 100; i < 1000; i++ {
		if isFlower(i) == true {
			fmt.Println(i)
		} else {
			continue
		}
	}
}

// TODO shui xian hua shu
func isFlower(n int) bool {
	var i int //bai wei
	var j int //shi wei
	var k int //ge wei

	i = int(n / 100)
	j = int((n % 100) / 10)
	k = n % 10
	if n == (i*i*i + j*j*j + k*k*k) {
		return true
	} else {
		return false
	}
}
