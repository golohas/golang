package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// TODO su shu ji suan

	fmt.Println("101~200 de su shu you :")
	startTime1 := time.Now().UnixNano()
	//fmt.Println(startTime1)
	for i := 101; i < 201; i++ {
		if isPrimeNormalOne(i) == true {
			fmt.Println(i)
		}
	}
	endTime1 := time.Now().UnixNano()
	//fmt.Println(endTime1)
	waste1 := endTime1 - startTime1
	fmt.Println("isPrimeNormalOne waste time is:", waste1)

	startTime2 := time.Now().Unix()
	for i := 101; i < 201; i++ {
		if isPrimeNormalTwo(i) == true {
			fmt.Println(i)
		}
	}
	endTime2 := time.Now().Unix()
	waste2 := endTime2 - startTime2
	fmt.Println("isPrimeNormalTwo waste time is:", waste2)
}

func isPrimeNormalOne(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeNormalTwo(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n%6 != 1 && n%6 != 5 {
		return false
	}

	tmp := int64(math.Sqrt(float64(n)))
	for i := 5; i <= int(tmp); i++ {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true

}
