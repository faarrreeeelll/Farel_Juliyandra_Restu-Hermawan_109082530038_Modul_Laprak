package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	CetakFaktor(1,n)
}

func CetakFaktor (i, n int) {
	if i > n {
		return
	}
	
	if n % i == 0 {
		fmt.Printf("%d ", i)
	}
	CetakFaktor(i + 1, n)
}
