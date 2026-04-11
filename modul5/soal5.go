package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	CetakDeret(1,n)
}

func CetakDeret (i, n int) {
	if i > n {
		return
	}
	CetakHasil(i)
	CetakDeret(i + 1, n)
}

func CetakHasil (i int) {
	if i % 2 != 0 {
		fmt.Printf("%d ", i)
	}
}
