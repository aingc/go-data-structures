package main

import "fmt"

// Stack holds a slice, LIFO
type Stack struct {
	items []int
}

// Queue holds a slice, FIFO
type Queue struct {
	items []int
}

// Push will add a value at the top of the stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove the value at the top of stack and return the removed value
func (s *Stack) Pop() int {
	stackTopIndex := len(s.items)-1
	removedVal := s.items[stackTopIndex]
	s.items = s.items[:stackTopIndex]
	return removedVal
}

// Enqueue will add a value at the end of the queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue will remove a value at the beginning of the queue and return the removed value
func (q *Queue) Dequeue() int {
	removedVal := q.items[0]
	q.items = q.items[1:len(q.items)]
	return removedVal
}


func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}