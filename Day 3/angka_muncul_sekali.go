package main

import (
	"fmt"
	"strconv"
)


func munculSekali(angka string) []int {
	unique := make(map[int]int, 0)
	resArr := make([]int, 0)
	for _, elem := range angka {
		conv, _ := strconv.Atoi(string(elem))
		unique[conv] += 1
	}
	for _, elem := range angka {
		conv, _ := strconv.Atoi(string(elem))
		if unique[int(conv)] == 1 {
			resArr = append(resArr, conv)
		}
	}
	return resArr
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
