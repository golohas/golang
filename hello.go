package main

import (
	"fmt"
)

func add(a int, b int, c chan int) {
	var sum int
	sum = a + b
	c <- sum
}

func main() {
	var pipe chan int
	/* fmt.Println("hello")
	fmt.Println("test println")
	*/
	/* for i := 0; i <= 100; i++ {
		go testPrint(i)
	} */
	// test_pipe()
	pipe = make(chan int, 2)
	// go add(2, 3, pipe)
	go add(3, 4, pipe)
	sum := <-pipe

	fmt.Println(sum)
}
