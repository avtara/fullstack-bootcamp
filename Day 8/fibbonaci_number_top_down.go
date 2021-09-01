package main

import "fmt"

func fibo(n int) int {
	a := make([]int, n)
	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		a[i-1] = fibMemo(i, m)
	}
	if len(a) == 0 {
		return 0
	}else {
		return a[n-1]
	}
}

func fibMemo(x int, m map[int]int) int {
	if x < 2 {
		m[x] = x
		return x
	}
	_, ok := m[x-1]
	if !ok {
		m[x-1] = fibMemo(x-1, m)
	}
	_, ok = m[x-2]
	if !ok {
		m[x-2] = fibMemo(x-2, m)
	}
	return m[x-1] + m[x-2]
}



func main() {
	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 13
	fmt.Println(fibo(10)) // 55
}
