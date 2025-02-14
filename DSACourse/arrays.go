package main

// import "strconv"
// import "encoding/json"
import (
	"fmt"
	"math"
)

func MaxMinArr(arr []int, size int) {
	//Implement your solution here
	last := 0
	for i := 0; i < size-1; i = i + 2 {
		last = arr[size-1]
		fmt.Println(arr)
		for j := size - 1; j > i; j-- {
			arr[j] = arr[j-1]
		}
		arr[i] = last
	}

}

func ArrayIndexMaxDiff(arr []int, size int) int {
	//Implement your solution here
	max := -1
	diff := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if arr[j] > arr[i] {
				if j > i {
					diff = j - i
					diff = int(math.Abs(float64(diff)))
					if diff > max {
						max = diff
					}
				}
			}
		}

	}
	return max
}

func maxPathSum(arr1 []int, size1 int, arr2 []int, size2 int) int {

	i := 0
	j := 0

	sum1 := 0
	sum2 := 0

	m := 0

	for i < size1 && j < size2 {
		if arr1[i] < arr2[j] {
			sum1 += arr1[i]
			i++
		} else if arr1[i] > arr2[j] {
			sum2 += arr2[j]
			j++
		} else {
			m += int(math.Max(float64(sum1), float64(sum2))) + arr1[i]
			sum1 = 0
			sum2 = 0
			i++
			j++
		}
	}

	for i < size1 {
		sum1 += arr1[i]
		i++
	}

	for j < size2 {
		sum2 += arr2[j]
		j++
	}

	m += int(math.Max(float64(sum1), float64(sum2)))

	return m
}

// func main() {
// 	// arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	// size := len(arr)
// 	// MaxMinArr(arr, size)
// 	// fmt.Println(arr)

// 	//

// 	arr1 := []int{12, 13, 18, 20, 22, 26, 70}
// 	size1 := 7
// 	arr2 := []int{11, 15, 18, 19, 20, 26, 30, 31}
// 	size2 := 8

// 	fmt.Println(maxPathSum(arr1, size1, arr2, size2))

// }
