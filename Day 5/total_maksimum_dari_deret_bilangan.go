package main

import "fmt"

func MaxSequence(arr []int) int {
	var res, maxLocal = arr[0], 0

	for idx, _ := range arr{
		maxLocal += arr[idx]
		if maxLocal < 0 {
			maxLocal = 0
		}else if res < maxLocal {
			res = maxLocal
		}
	}
	return res
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12

}