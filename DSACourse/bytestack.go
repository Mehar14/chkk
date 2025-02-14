package main

import "fmt"

type StackByte struct {
	s []byte
}

func (s *StackByte) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *StackByte) Length() int {
	length := len(s.s)
	return length
}

func (s *StackByte) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func (s *StackByte) Push(value byte) {
	s.s = append(s.s, value)
}

func (s *StackByte) Pop() byte {
	if s.IsEmpty() == true {
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *StackByte) Top() byte {
	if s.IsEmpty() == true {
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}

// func maxDepthParenthesis(expn string, size int) int {
// 	//Implement your solution here

// 	//Uncomment the ch variable and use it to move in an expression array
// 	//var ch byte
// 	max := 0
// 	stk := StackByte{}
// 	for _, ch := range expn {
// 		if ch == '(' {
// 			stk.Push(byte(ch))
// 		} else if ch == ')' {
// 			stk.Pop()
// 		}

// 		if stk.Length() > max {
// 			max = stk.Length()
// 		}
// 	}

// 	return max //Return 0 if depth of parenthesis is zero
// }

// func main() {
// 	exp := ""
// 	fmt.Println(longestContBalParen(exp, len(exp)))
// }
