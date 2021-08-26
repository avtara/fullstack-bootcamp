package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		return pow(x, n / 2) * pow(x, n / 2)
	} else {
		return x * pow(x, n - 1)
	}
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(7, 2))  // 49
	fmt.Println(pow(10, 5)) // 100000
	fmt.Println(pow(17, 6)) // 24137569
	fmt.Println(pow(5, 3))  // 125
}

/*

	pow(5,1) = 5 * 1
	pow(5,2) = pow(5,1) * pow(5,1) = (5 * 1) * (5 * 1) = 25
	pow(5,3) = 5 * pow(5,2) = 5 * pow(5,2) = 5 * pow(5,1) * pow(5,1) = 5 * (5 * 1) * (5 * 1) = 75

*/
