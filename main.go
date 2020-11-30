package main

import (
	"fmt"
	"queue/queue"
)

func main() {
	fmt.Println("Hello World")

	q := queue.New()

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(1)

	q.Print()

	fmt.Printf("popped: %v\n", q.Pop())

	q.Print()
}

// push 1,2,3,1 -> 2,3,1
