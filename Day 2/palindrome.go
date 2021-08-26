package main

import "fmt"



func palindrome(input string) bool {
	var last_index = len(input) - 1
	for i := 0; i < len(input); i++ {
		if input[i] != input[last_index]{
			return false
		}
		last_index--
	}
	return true
}



func main() {
	fmt.Println(palindrome("civic"))       // true
	fmt.Println(palindrome("katak"))       // true
	fmt.Println(palindrome("mistar"))      // false
	fmt.Println(palindrome("lion"))        // false
	fmt.Println(palindrome("kasur rusak")) // true
}