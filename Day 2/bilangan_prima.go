package main

import "fmt"

func primeNumber(number int) bool {
	if number > 1 {
		for i := 2; i < number; i++{
			if number % i == 0{
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(primeNumber(11)) // true
	fmt.Println(primeNumber(13)) // true
	fmt.Println(primeNumber(17)) // true
	fmt.Println(primeNumber(20)) // false
	fmt.Println(primeNumber(35)) // false
	//var data int
	//fmt.Scanf("%d", &data)
	//fmt.Println(primeNumber(data))
}