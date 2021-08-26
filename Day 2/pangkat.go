package main

import "fmt"

func pangkat(base, pangkat int) int {
	var result int = base
	for i := 1; i < pangkat; i++{
		result *= base
	}
	return result
}



func main() {
	fmt.Println(pangkat(2, 3))  // 8
	fmt.Println(pangkat(5, 3))  // 125
	fmt.Println(pangkat(10, 2)) // 100
	fmt.Println(pangkat(2, 5))  // 32
	fmt.Println(pangkat(7, 3))  // 343

}