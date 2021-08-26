package main

import(
	"fmt"
)

func main() {
	var T, r float64
	fmt.Scanf("%f %f", &T, &r)
	const PHI float64 = 3.14
	var result float64 = 2 * PHI * r * (r + T)
	fmt.Println(result)
}