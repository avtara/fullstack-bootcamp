
package main


import (
	"fmt"
	"math"
)

func primeX(number int) int {
	var  i,num,res int = 0,1,0
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

	for i < number {
		if check(num) {
			i++
			res = num
		}
		num += 1
	}
	return res
}



func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29

}