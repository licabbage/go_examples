package main

import "fmt"

// 在一个channel关闭后，仍然可以通过range 来访问其元素
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
