package main

import (
	"fmt"
)

func main() {
	miList := []int{1, 3, 7, 4, 2}
	BubbleSort(miList)
	fmt.Print(miList)
}

func BubbleSort(intList []int) {
	for i := 0; i < len(intList)-1; i++ {
		swapped := false
		for j := 0; j < len(intList)-i-1; j++ {
			if intList[j] > intList[j+1] {
				intList[j], intList[j+1] = intList[j+1], intList[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
