package main

import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	CetakDeret(1, n)
}

func CetakDeret(i, n int){
	if i <= n {
		angkaSekarang := n - i + 1
		fmt.Printf("%d ", angkaSekarang)
		if i < n {
			CetakDeret(i + 1, n)
			fmt.Printf("%d ", angkaSekarang)
		}
	}

}

