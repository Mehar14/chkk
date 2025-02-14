package main

import (
	"fmt"
	"strconv"
)

// type StackInt struct {
// 	data []int
// }

// func (s *StackInt) Length() int {
// 	return len(s.data)
// }
// func (s *StackInt) IsEmpty() bool {
// 	return s.Length() == 0
// }

// func (s *StackInt) Print() {
// 	fmt.Println(s.data)
// }

// func (s *StackInt) Push(value int) {
// 	s.data = append(s.data, value)
// }

// func (s *StackInt) Pop() int {
// 	if s.IsEmpty() {
// 		fmt.Println("The stack is empty.")
// 		return -9999
// 	} else {
// 		topind := s.Length() - 1
// 		element := s.data[topind]
// 		s.data = s.data[:topind]

// 		return element
// 	}
// }

// func (s *StackInt) Top() int {
// 	if s.IsEmpty() {
// 		fmt.Println("The stack is empty.")
// 		return -9999
// 	} else {
// 		topind := s.Length() - 1
// 		element := s.data[topind]
// 		return element
// 	}
// }

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int
}

func (s *StackLinkedList) Size() int {
	return s.size
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.Size() == 0
}

func (s *StackLinkedList) Peek() (bool, int) {
	if s.IsEmpty() {
		return true, 0
	} else {
		return false, s.head.value
	}
}
func (s *StackLinkedList) Push(val int) {

	node := &Node{}
	node.value = val
	node.next = s.head
	s.head = node
	s.size++

}
func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		return -9999, true
	} else {
		element := s.head.value
		s.head = s.head.next

		s.size--

		return element, false
	}
}

func (s *StackLinkedList) Print() {

	current := s.head

	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}

}

func sortedInsert(stk *StackInt, element int) {
	temp := StackInt{}

	for stk.Top() > element {
		popped := stk.Pop()
		temp.Push(popped)
	}

	stk.Push(element)

	for temp.Length() > 0 {
		popped := temp.Pop()
		stk.Push(popped)
	}
}

func sortedInsertRecursion(stk *StackInt, element int) {
	if stk.Length() == 0 || element > stk.Top() {
		stk.Push(element)
	} else {
		temp := stk.Pop()
		sortedInsertRecursion(stk, element)
		stk.Push(temp)
	}
}

func sortStack(stk *StackInt) {
	if stk.Length() == 0 {
		return
	} else {
		element := stk.Pop()
		sortStack(stk)
		sortedInsertRecursion(stk, element)
	}
}

func bottomInsert(stk *StackInt, element int) {
	//Implement your solution here
	temp := StackInt{}
	for stk.Length() > 0 {
		temp.Push(stk.Pop())
	}
	stk.Push(element)
	for temp.Length() > 0 {
		stk.Push(temp.Pop())
	}
}

func reverseStack(stk *StackInt) {
	//Implement your solution here
	if stk.Length() == 0 {
		return
	} else {
		element := stk.Pop()
		reverseStack(stk)
		bottomInsert(stk, element)
	}

}

func priority(a int) int {
	if a == 40 || a == 41 {
		return 0
	} else if a == 94 {
		return 5
	} else if a == 47 || a == 42 {
		return 3
	} else if a == 43 || a == 45 {
		return 2
	} else {
		return -1 //if not an operator
	}
}

func symbol(b int) string {
	if b == 94 {
		return "^"
	} else if b == 40 {
		return "("
	} else if b == 41 {
		return ")"
	} else if b == 42 {
		return "*"
	} else if b == 47 {
		return "/"
	} else if b == 43 {
		return "+"
	} else if b == 45 {
		return "-"
	}
	return ""
}

func InfixToPostfix(expn string) string {

	output := ""
	stk := StackInt{}

	for c := 0; c < len(expn); c++ {
		fmt.Print(output)
		stk.Print()
		// Use the variable c here
		prio := priority(int(expn[c]))
		if prio == -1 {
			//append to the output
			output = output + string(expn[c])
		} else {
			//is an operator
			//check if stack empty
			if stk.Length() == 0 {
				stk.Push(int(expn[c]))
			} else {
				//stack is not empty
				//check the top of stack
				if priority(stk.Top()) < prio {
					stk.Push(int(expn[c]))
				} else if symbol(int(expn[c])) == "(" {
					stk.Push(int(expn[c]))
				} else if symbol(int(expn[c])) == ")" {

					for stk.Length() > 0 && symbol(stk.Top()) != "(" {
						temp := stk.Pop()
						output = output + symbol(temp)
					}
					stk.Pop()
				} else {
					for stk.Length() > 0 && priority(stk.Top()) >= prio {
						temp := stk.Pop()
						output = output + symbol(temp)
					}

					stk.Push(int(expn[c]))
				}
			}
		}

	}

	for stk.Length() > 0 {
		temp := stk.Pop()
		output = output + symbol(temp)

	}

	//Implement your solution here

	return output
}

func InfixToPrefix(expn string) string {
	//Implement your solution here

	exp := ""

	for c := len(expn) - 1; c >= 0; c-- {

		if string(expn[c]) == ")" {
			exp = exp + "("
		} else if string(expn[c]) == "(" {
			exp = exp + ")"
		} else {
			exp = exp + string(expn[c])
		}

	}

	postfix := InfixToPostfix(exp)

	prefix := ""

	for c := len(postfix) - 1; c >= 0; c-- {
		prefix = prefix + string(postfix[c])
	}

	//kindly replace the string with your desired variable or string
	return prefix
}

type MinStack struct {
	maxSize int
	stk     StackInt
	temp    StackInt
	// Initialize your data structures here
}

// removes and returns value from stack
func (m *MinStack) Pop() int {
	// write your code here
	if m.stk.Length() == 0 {
		return -9999
	} else {
		topind := m.stk.Length() - 1
		top := m.stk.data[topind]
		return top
	}
}

// Pushes value into the stack
func (m *MinStack) Push(value int) {
	// write your code here
	if m.stk.Length() == 0 {
		m.stk.Push(value)
	} else {

		for m.stk.Top() < value {
			popped := m.stk.Pop()
			m.temp.Push(popped)
		}

		m.stk.Push(value)

		for m.temp.Length() > 0 {
			popped := m.temp.Pop()
			m.stk.Push(popped)
		}
	}

}

//()(())(())())

// func longestContuBalParen(str string, _ int) int {
// 	length := 0
// 	stk := StackInt{}
// 	stk.Push(-1)

// 	for i, ch := range str {
// 		stk.Print()
// 		if ch == '(' {
// 			stk.Push(i)
// 		} else if ch == ')' {
// 			stk.Pop()

// 			if stk.Length() > 0 {
// 				if i-stk.Top() > length {
// 					length = i - stk.Top()
// 				}
// 			} else {
// 				stk.Push(i)
// 			}

// 		}
// 	}

// 	return length
// }

// returns minimum value in O(1)
func (m *MinStack) Min() int {
	// write your code here
	return m.stk.Top()
}

func findDuplicateParenthesis(expn string, size int) bool {
	stk := StackInt{}

	dupl := false
	for i := 0; i < size; i++ {

		if string(expn[i]) != ")" {
			stk.Push(i)
		} else {
			count := 0
			for string(expn[stk.Top()]) != "(" {
				stk.Pop()
				count++
			}
			stk.Pop()
			if count <= 1 {
				dupl = true
			}

		}

	}
	return dupl
}

func printParenthesisNumber(expn string, size int) {
	//Uncomment the ch variable and use it iterate through the string
	var ch byte
	output := "" //use output variable to save and print the output string

	stk := StackInt{}

	//Implement your solution here
	count := 1
	for i := 0; i < size; i++ {
		stk.Print()
		ch = expn[i]

		if ch == '(' {
			stk.Push(count)
			output = output + strconv.Itoa(count)
			count++
		} else if ch == ')' {
			output = output + strconv.Itoa(stk.Pop())
		}
	}
	fmt.Println(output)
}

// func reverseKElementInStack(stk *StackInt, k int) {
// 	//Implement your solution here
// 	que := new(Queue)
// 	for k > 0 && stk.Length() > 0 {
// 		element := stk.Pop()
// 		que.Enqueue(element)
// 		k--
// 	}
// 	for que.Length() > 0 {
// 		stk.Push(que.Dequeue.(int))
// 	}

// }

// func main() {
// 	exp := "(((a+(b))+(c+d)))"
// 	printParenthesisNumber(exp, len(exp))

// 	// 	exp := "()(())(())())"

// 	// 	fmt.Println(longestContuBalParen(exp, len(exp)))

// 	// 	// 	mins := MinStack{}
// 	// 	// 	mins.Push(10)
// 	// 	// 	mins.Push(2)
// 	// 	// 	mins.Push(4)

// 	// 	// 	fmt.Println(mins.Min())

// 	// 	// 	// 	// stack := StackInt{}

// 	// 	// 	// 	// stack.Push(10)
// 	// 	// 	// 	// stack.Push(19)
// 	// 	// 	// 	// stack.Push(20)

// 	// 	// 	// 	// stack.Print()

// 	// 	// 	// 	// fmt.Println("Element at the top is", stack.Top())

// 	// 	// 	// 	// ele := stack.Pop()
// 	// 	// 	// 	// fmt.Println(ele)

// 	// 	// 	// 	// stack.Print()

// 	// 	// 	// 	// s := new(StackLinkedList)

// 	// 	// 	// 	// fmt.Println(s.head)
// 	// 	// 	// 	// s.Push(1)
// 	// 	// 	// 	// fmt.Println(s.head)

// 	// 	// 	// 	// s.Push(2)
// 	// 	// 	// 	// s.Push(3)
// 	// 	// 	// 	// _, val := s.Peek()
// 	// 	// 	// 	// fmt.Println("Peek() of a stack is:", val)
// 	// 	// 	// 	// fmt.Print("Stack consist following elements: ")
// 	// 	// 	// 	// for s.IsEmpty() == false {
// 	// 	// 	// 	// 	val, _ := s.Pop()
// 	// 	// 	// 	// 	fmt.Print(val, " ")
// 	// 	// 	// 	// }
// 	// 	// 	// 	// stk := new(StackInt)
// 	// 	// 	// 	// stk.Push(1)
// 	// 	// 	// 	// stk.Push(2)
// 	// 	// 	// 	// stk.Push(9)
// 	// 	// 	// 	// stk.Push(4)
// 	// 	// 	// 	// // sortedInsertRecursion(stk, 3)
// 	// 	// 	// 	// // sortedInsertRecursion(stk, 7)
// 	// 	// 	// 	// // sortedInsertRecursion(stk, 8)
// 	// 	// 	// 	// // stk.Print()
// 	// 	// 	// 	// stk.Print()
// 	// 	// 	// 	// sortStack(stk)
// 	// 	// 	// 	// stk.Print()

// 	// 	// 	// 	// bottomInsert(stk, 5)
// 	// 	// 	// 	// stk.Print()

// 	// 	// 	// 	// reverseStack(stk)
// 	// 	// 	// 	// stk.Print()

// 	// 	// 	// 	// exp := "{[({})]}"

// 	// 	// 	// exp := "a+b*((c^d-e))^(f+g*h)-i"

// 	// 	// 	// fmt.Println(InfixToPostfix(exp))

// 	// 	// 	// fmt.Println(InfixToPrefix(exp))

// 	// 	// 	// h := "()*/+-"
// 	// 	// 	// for c := 0; c < len(h); c++ {
// 	// 	// 	// 	fmt.Print(string(h[c]))
// 	// 	// 	// 	fmt.Println(int(h[c]))

// 	// 	// 	// }

// }

// // abcd^e-fgh*+^*+i-
