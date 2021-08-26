package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var resArr = make([]int, 0)
	var hash = make(map[int]int, 0)
	for idx, elem:= range arr {
		minusTarget := target - elem
		if pos, ok := hash[minusTarget]; ok {
			if pos != idx {
				resArr = append(resArr, hash[minusTarget], idx)
			}
		} else {
			hash[elem] = idx
		}
	}
	return resArr
}



func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
