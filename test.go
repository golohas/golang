package main

import (
	"fmt"
)

func testPrint(a int) {
	fmt.Println(a)
}

func test_pipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2

	fmt.Println(len(pipe))
}

func test() {
	fmt.Println("test")
}
