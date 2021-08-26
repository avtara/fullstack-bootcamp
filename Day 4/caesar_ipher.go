package main

import (
	"fmt"
)



func caesar(offset int, input string) string {
	var res string
	for i := 0; i < len(input); i++ {
		if  input[i] + uint8(offset % 26) > 122{
			res += string(input[i] + uint8(offset % 26) - 26)
		}else {
			res += string(input[i] + uint8(offset % 26))
		}
	}
	return res
}



func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) 	// bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))	// mnopqrstuvwxyzabcdefghijkl
}
