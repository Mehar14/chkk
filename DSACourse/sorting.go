package main

import (
	"fmt"
)

// import (
// 	"fmt"
// 	"strconv"
// )

func BubbleSort(arr []int) {
	size := len(arr)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				//swap
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}

		fmt.Println(arr)
	}

}

func modifiedBubbleSort(arr []int) {
	size := len(arr)
	swapped := 1

	for i := 0; i < size; i++ {
		swapped = 1
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				//swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = 0
			}
		}

		fmt.Println(arr)

		if swapped == 1 {
			break
		}
	}
}

func insertionSort(arr []int) {
	size := len(arr)

	for i := 1; i < size; i++ {
		for j := 0; j <= i-1; j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}

			fmt.Println(arr)
		}

		//fmt.Println(arr)
	}
}

func selectionsort(arr []int) {
	size := len(arr)

	for i := 0; i < size; i++ {
		max := arr[0]
		max_ind := 0
		for j := 0; j < size-i; j++ {
			if arr[j] > max {
				max = arr[j]
				max_ind = j
			}
		}
		temp := arr[size-i-1]
		arr[size-i-1] = arr[max_ind]
		arr[max_ind] = temp

		fmt.Println(arr)
	}
}

func mergesort(arr []int) []int {
	fmt.Println(arr)
	if len(arr) <= 1 {
		return arr
	} else {
		mid := len(arr) / 2
		left := mergesort(arr[:mid])
		right := mergesort(arr[mid:])
		return merge(left, right)
	}

}

func merge(left []int, right []int) []int {

	var result []int

	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i = i + 1
		} else {
			result = append(result, right[j])
			j = j + 1
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}

func mymergesort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2

	left := mymergesort(arr[:mid])
	right := mymergesort(arr[mid:])

	return mymerge(left, right)
}

func mymerge(left []int, right []int) []int {
	i := 0
	j := 0

	var result []int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i = i + 1
		} else {
			result = append(result, right[j])
			j = j + 1
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func quicksort(arr []int) {
	quicksortutil(arr, 0, len(arr)-1)
}

func quicksortutil(arr []int, lower int, upper int) {
	if upper <= lower {
		return
	}

	pivot := arr[lower]
	start := lower
	stop := upper

	for lower < upper {
		for arr[lower] <= pivot && lower < upper {
			lower = lower + 1
		}

		for arr[upper] > pivot && lower <= upper {
			upper = upper - 1
		}

		if lower < upper {
			arr[lower], arr[upper] = arr[upper], arr[lower]
		}

	}
	arr[start], arr[upper] = arr[upper], arr[start]
	quicksortutil(arr, start, upper-1)
	quicksortutil(arr, upper+1, stop)

}

func myquicksort(arr []int, lower int, upper int) {
	fmt.Println(arr)
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper

	fmt.Println(pivot)

	for lower < upper {
		for pivot >= arr[lower] && lower < upper {
			lower = lower + 1
		}

		for pivot < arr[upper] && lower <= upper {
			upper = upper - 1
		}

		if lower < upper {
			arr[lower], arr[upper] = arr[upper], arr[lower]

		}
	}

	arr[start], arr[upper] = arr[upper], arr[start]
	myquicksort(arr, start, upper-1)
	myquicksort(arr, upper+1, stop)

}

func quickselect(arr []int, lower int, upper int, k int) int {
	fmt.Println(arr)
	if k > 0 && k <= upper-lower+1 {
		pivot := arr[upper]
		start := lower
		stop := lower

		for stop <= upper-1 {
			if arr[stop] <= pivot {
				arr[stop], arr[start] = arr[start], arr[stop]
				start++
			}
			stop++
		}

		arr[start], arr[upper] = arr[upper], arr[start]
		pivotInd := start

		if pivotInd-lower == k-1 {
			return arr[pivotInd]
		}
		if pivotInd-lower > k-1 {
			return quickselect(arr, lower, pivotInd-1, k)
		} else {
			return quickselect(arr, pivotInd+1, upper, k-pivotInd+lower-1)
		}
	} else {
		return -1
	}

}

func countsort(arr []int, min int, max int) {
	rng := (max - min) + 1
	countarr := make([]int, rng)

	//fmt.Println(countarr)

	for i := 0; i < len(arr); i++ {
		//fmt.Println(i)
		countarr[arr[i]-1]++
	}

	//fmt.Println(countarr)
	k := 0
	for j := 0; j < len(countarr); j++ {
		for countarr[j] > 0 {
			arr[k] = j + 1
			k++
			countarr[j]--
		}

		//fmt.Println(arr)
	}

}

func MinMax(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min, max
}

func Partition01(arr []int, size int) int {
	//Write your code here

	start := 0
	stop := size - 1

	count := 0

	for start <= stop {

		for arr[start] == 0 {
			start++
		}

		for arr[stop] == 1 {
			stop--
		}

		if start < stop {
			arr[start], arr[stop] = arr[stop], arr[start]
			count = count + 1

		}
	}
	print(count)
	return count
}

func Partition012(arr []int, size int) {

	//Implement your solution here

	start := 0
	stop := size - 1
	i := 0

	for i <= stop {

		fmt.Println(arr)

		if arr[i] == 0 {
			arr[i], arr[start] = arr[start], arr[i]
			fmt.Println(start)
			i += 1
			start += 1
		} else if arr[i] == 2 {
			arr[i], arr[stop] = arr[stop], arr[i]
			stop -= 1
		} else {
			i += 1
		}

	}
}

// func main() {
// 	//data := []int{5, 6, 2, 4, 7, 3}
// 	//modifiedBubbleSort(data)

// 	//insertionSort((data)

// 	//selectionsort(data)

// 	//fmt.Println(mymergesort(data))

// 	//myquicksort(data, 0, len(data)-1)

// 	// thirdsmallest := quickselect(data, 0, len(data)-1, 3)

// 	// fmt.Println(thirdsmallest)

// 	// fmt.Println(data)

// 	data := []int{2, 6, 4, 1, 5, 8, 1, 4, 6, 1}

// 	min, max := MinMax(data)

// 	countsort(data, min, max)

// 	// fmt.Println(data)

// 	// seronon1 := []int{1, 0, 1, 1, 0, 0, 0, 0}
// 	// Partition01(seronon1, len(seronon1))

// 	// fmt.Println(seronon1)

// 	a := []int{0, 1, 0, 2, 0, 1, 0, 1}
// 	Partition012(a, len(a))

// 	fmt.Println(a)

// }
