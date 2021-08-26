package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var res string
	if len(b) > len(b){
		a,b = b,a
	}
	for i := 0; i < len(b); i++ {
		if strings.Contains(a,b[:i+1]){
			res = b[:i+1]
		}
	}
	return res
}



func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
