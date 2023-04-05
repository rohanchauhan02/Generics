package main

import "fmt"

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s=append(*s, v)
}
func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("stack is emptry")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func main() {
	myStack := make(Stack[any],0)
	fmt.Printf("STACK :%v\n", myStack)
	myStack.Push("rohan")
	myStack.Push([]int{1,2,3,4})
	myStack.Push("hello")
	fmt.Printf("STACK :%v\n", myStack)
	fmt.Printf("POP Item: %v ", myStack.Pop())
	fmt.Printf("STACK :%v\n", myStack)
}
