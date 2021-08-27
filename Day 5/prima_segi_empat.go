package main

import (
	"fmt"
	"math"
)

func primaSegiEmpat(high, wide, start int) {
	var total, i, j int
	start += 1

	check := func(data int) bool {
		if data < 2 {
			return false
		}

		sqrtNum := int(math.Sqrt(float64(data)))
		for i := 2; i <= sqrtNum; i++ {
			if data % i == 0 {
				return false
			}
		}
		return true
	}

	for i < wide {
		if check(start) {
			j++
			total += start
			fmt.Printf("%d\t", start)
		}
		if j == high {
			i++
			j = 0
			fmt.Println()
		}
		start += 1
	}
	fmt.Println(total)
}

func main() {
	primaSegiEmpat(2, 3, 13)
	/*
	   17 19
	   23 29
	   31 37
	   156
	*/
	primaSegiEmpat(5, 2, 1)

	/*
	   2  3  5  7 11
	   13 17 19 23 29
	   129
	*/
}