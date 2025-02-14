package main

import (
	"fmt"
	"strconv"
)

func Factorial(i int) int {
	// Implement your solution here
	if i == 0 {
		return 1
	} else {
		return i * Factorial(i-1)
	}
}

func printxInt(number int, count int) {
	// Write recursive code here
	if number < 16 {
		fmt.Print(count)
		if number == 15 {
			fmt.Print("F")
		} else if number == 14 {
			fmt.Print("E")
		} else if number == 13 {
			fmt.Print("D")
		} else if number == 12 {
			fmt.Print("C")
		} else if number == 11 {
			fmt.Print("B")
		} else if number == 10 {
			fmt.Print("A")
		} else {
			fmt.Print(number)
		}
	} else {
		printxInt(number-16, count+1)
	}
}

func printInt(number int) {
	// Write recursive code here
	remainder := number % 16
	if number != 0 {
		printInt(number / 16)
	}
	if remainder == 15 {
		fmt.Print("F")
	} else if remainder == 14 {
		fmt.Print("E")
	} else if remainder == 13 {
		fmt.Print("D")
	} else if remainder == 12 {
		fmt.Print("C")
	} else if remainder == 11 {
		fmt.Print("B")
	} else if remainder == 10 {
		fmt.Print("A")
	} else if number == 0 {
		// fmt.Print("A")
	} else {
		fmt.Print(remainder)
	}
	// fmt.Println("-1")
	// Do not remove the print statement although you can change it with your initialized variables.
}

// func GCD(m int, n int) int{
// 	smaleest:=-1
// 	if m < n{
// 		smaleest= m
// 	} else {
// 		smaleest= n
// 	}

// 	max:=1

// 	for i:=1;i<smaleest+1;i++{
// 		if m%i==0 && n%i == 0{

// 			if i > max{
// 				max = i
// 			}
// 		}
// 	}

// 	return max
// }

func GCD(m int, n int) int {
	if n == 0 {
		return m
	} else {
		return GCD(n, m%n)
	}
}

func Fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

func Permutation(data []int, i int, length int) {

	// base condition DO NOT ALTER IT
	if length == i {
		temp := "{"
		for k := 0; k < length; k++ {
			temp += strconv.Itoa(data[k])
			temp += " "
		}
		temp += "} "
		fmt.Print(temp)
		return
	} else {
		for j := i; j < length; j++ {
			data[i], data[j] = data[j], data[i]
			Permutation(data, i+1, length)
			data[i], data[j] = data[j], data[i]
		}
	}

}

func towerOfHanoi(num int, src byte, dst byte, temp byte) {
	// write some code here
	if num == 0 {
		return
	} else {
		//DONOT change print statement
		towerOfHanoi(num-1, src, temp, dst)
		fmt.Printf("Move %d disk from peg %c to peg %c ", num, src, dst)
		fmt.Println("")
		towerOfHanoi(num-1, temp, dst, src)
		// write some code here
	}

}

func BinarySearchRecursive(data []int, value int) bool {
	size := len(data)
	return BinarySearchRecursiveUtil(data, 0, size-1, value)
}

func BinarySearchRecursiveUtil(data []int, low int, high int, value int) bool {
	//Write your code here
	if low > high {
		return false
	}
	midval := (low + high) / 2
	if data[midval] == value {

		return true
	} else if value < data[midval] {
		return BinarySearchRecursiveUtil(data, low, midval-1, value)
	} else {
		return BinarySearchRecursiveUtil(data, midval+1, high, value)
	}
}

// func main() {
// 	//fmt.Println(Factorial(5))
// 	//printInt(95)
// 	//fmt.Println(GCD(18,24))

// 	//fmt.Println(Fibonacci(5))
// 	// data:= []int{0,1,2}
// 	// Permutation(data,0,3)

// 	// towerOfHanoi(5,'A','C','B')
// 	arr1 := []int{1, 3, 4, 6, 8, 10}
// 	value := 7
// 	fmt.Println(BinarySearchRecursive(arr1, value))

// }
