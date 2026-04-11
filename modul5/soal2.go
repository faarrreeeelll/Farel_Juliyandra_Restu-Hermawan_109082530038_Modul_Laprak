package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	CetakDeret(1,n)
}

func CetakDeret (i, n int) {
	if i <= n {
		CetakBintang(1,i)
		fmt.Println()
		CetakDeret(i + 1, n)
	}
}

func CetakBintang (kolom, jumlah int) {
	if kolom <= jumlah {
		fmt.Print("*")
		CetakBintang(kolom + 1, jumlah)
	}	
}