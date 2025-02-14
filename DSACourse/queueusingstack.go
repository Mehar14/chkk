package main

import (
	"fmt"
)

type StackQueue struct {
	s []int
}

func (s *StackQueue) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *StackQueue) Length() int {
	length := len(s.s)
	return length
}

func (s *StackQueue) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func (s *StackQueue) Push(value int) {
	s.s = append(s.s, value)
}

func (s *StackQueue) Pop() int {
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *StackQueue) Top() int {
	length := len(s.s)
	res := s.s[length-1]
	return res
}

type QueueUsingStack struct {
	stk1 StackQueue
	stk2 StackQueue
	// add more values here if need be

}

func (que *QueueUsingStack) Add(value int) {
	//Implement your solution here
	if que.stk1.IsEmpty() {
		que.stk1.Push(value)
	} else {
		for que.stk1.Length() > 0 {
			temp := que.stk1.Pop()
			que.stk2.Push(temp)
		}

		que.stk1.Push(value)

		for que.stk2.Length() > 0 {
			temp := que.stk2.Pop()
			que.stk1.Push(temp)
		}
	}
}

func (que *QueueUsingStack) Remove() int {
	//Implement your solution here

	if que.stk1.IsEmpty() {
		return -9999
	}

	return que.stk1.Pop()
}

func (que *QueueUsingStack) Length() int {
	//Implement your solution here

	return que.stk1.Length()
}

func (que *QueueUsingStack) IsEmpty() bool {
	//Implement your solution here

	return que.stk1.IsEmpty()
}

func reverseQueue(que *Queue) {
	//Implement your solution here
	stk := StackInt{}
	for que.Length() > 0 {
		stk.Push(que.Dequeue().(int))
	}

	for stk.Length() > 0 {
		que.Enqueue(stk.Pop())
	}

}

func reverseKElementInQueue(que *Queue, k int) {
	//Implement your solution here
	stk := StackInt{}
	que2 := Queue{}

	for {
		stk.Push(que.Dequeue().(int))
		k--

		if k == 0 {
			break
		}

	}

	for stk.Length() > 0 {
		que2.Enqueue(stk.Pop())
	}

	for que.Length() > 0 {
		que2.Enqueue(que.Dequeue().(int))
	}

	for que2.Length() > 0 {
		que.Enqueue(que2.Dequeue().(int))
	}

}

// func main() {

// }
