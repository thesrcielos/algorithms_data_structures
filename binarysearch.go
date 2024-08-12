package main

import (
	"fmt"
)

func main() {
	integerList := []int{-1, 2, 4, 6, 7, 8, 9, 12}
	fmt.Print(binarySearch(integerList, 2))

}

func binarySearch(intList []int, target int) int {
	low := 0
	high := len(intList) - 1
	for low <= high {
		m := low + (high-low)/2
		value := intList[m]
		if value == target {
			return m
		}
		if value > target {
			high = m + 1
		} else {
			low = m + 1
		}
	}
	return -1
}
