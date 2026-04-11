package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	CetakDeret(0,n)
}

func CetakDeret (i, n int) {
	if i <= n {
		fmt.Printf("%d ",Fibonacci(i))
		CetakDeret(i + 1, n)
	}
}

func Fibonacci(i int) int {
	if i <= 0 {
		return 0
	}else if i == 1 {
		return 1
	}else {
		return Fibonacci(i-1) + Fibonacci(i-2)
	}
}