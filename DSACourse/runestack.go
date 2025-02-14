package main

import (
	"fmt"
	"math"
)

type Stack struct {
	s []rune
}

func (s *Stack) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *Stack) Length() int {
	length := len(s.s)
	return length
}

func (s *Stack) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func (s *Stack) Push(value rune) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() == true {
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Top() rune {
	if s.IsEmpty() == true {
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}

func IsBalancedParenthesis(expn string) bool {
	//Implement your solution here
	ind := -1
	stk := Stack{}

	opn := []rune{'{', '[', '('}
	cls := []rune{'}', ']', ')'}

	for i := 0; i < len(cls); i++ {
		if rune(expn[0]) == cls[i] {
			return false
		}
	}

	for c := range expn {
		for i := 0; i < len(opn); i++ {

			if rune(expn[c]) == opn[i] {
				top := stk.Top()

				for j := 0; j < len(opn); j++ {
					if top == opn[j] {
						ind = j

					}
				}

				if ind > i {
					return false
				}

				stk.Push(rune(expn[c]))

			} else if rune(expn[c]) == cls[i] && stk.Length() > 0 {
				top := stk.Top()

				for j := 0; j < len(opn); j++ {
					if top == opn[j] {
						ind = j

					}
				}

				if ind == i {
					stk.Pop()
				} else {
					return false
				}
			}
		}
	}

	if stk.Length() == 0 {
		return true
	} else {
		return false
	}
}

func reverseParenthesis(expn string, size int) int {
	new := ""
	stk := Stack{}

	closing := 0
	opening := 0
	var ch rune
	if size%2 != 0 {
		return -1
	} else {
		for i := 0; i < size; i++ {
			ch = rune(expn[i])

			if ch == '(' {
				stk.Push(ch)
			} else {
				if stk.Length() > 0 {
					stk.Pop()

				} else {
					new = new + string(expn[i])
					closing++
				}
			}
		}

		for stk.Length() > 0 {
			new = new + string(stk.Pop())
			opening++
		}

	}

	fmt.Println(new)
	fmt.Println("opening", opening)
	fmt.Println("closing", closing)
	reversal := int(math.Ceil(float64(opening)/2.0)) + int(math.Ceil(float64(closing)/2.0))
	//Implement your solution here

	//Uncomment this line of code and use ch to iterate in a string

	return reversal //Return 0 if string is empty
}

// func main() {
// 	exp := ")(())((("

// 	fmt.Println(reverseParenthesis(exp, len(exp)))
// 	// fmt.Println(IsBalancedParenthesis(exp))

// }
