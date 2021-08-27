package main

import (
	"fmt"
)

func FindMinAndMax(arr []int) string {
		max,min := arr[0], arr[0]
		var maxIdx, minIdx int
		for idx, val := range arr{
			if max < val {
				max = val
				maxIdx = idx
			}else if min > val {
				min = val
				minIdx = idx
			}
		}
		return fmt.Sprintf("min: %v index: %v max: %v index: %v", min, minIdx, max, maxIdx)
}



func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	// min: -2 index: 3 max: 8 index: 5
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	// min: -5 index: 1 max: 22 index: 3
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	// min: -21 index: 4 max: 9 index: 2
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	// min: -1 index: 0 max: 18 index: 5
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
	// min: -20 index: 5 max: 7 index: 4

}